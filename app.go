package cola

import (
	"bytes"
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zedisdog/cola/cmd/cola/migrate"
	"github.com/zedisdog/cola/database/seeder"
	"github.com/zedisdog/cola/task"
	"github.com/zedisdog/cola/transport/http"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"strings"
)

// Config 配置
//  config: 配置文件内容的bytes,yml格式
//  migration: 迁移文件的虚拟文件系统
func Config(config []byte) {
	// 读config
	viper.SetConfigType("yml")
	err := viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		panic(err)
	}

	// 环境变量中的__换成. 然后读环境变量，覆盖到配置
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	// 设置内置配置
	// 设置根
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.Set("root", wd)
}

type Server struct {
	logger     *logrus.Logger
	conf       *viper.Viper
	queue      *task.Queue
	httpServer *http.Server
	route      *gin.Engine
	migration  embed.FS
	db         *gorm.DB
}

func NewServer(config *viper.Viper, route *gin.Engine, db *gorm.DB, logger *logrus.Logger, migration embed.FS) *Server {
	s := &Server{
		logger:    logger,
		migration: migration,
		db:        db,
		conf:      config,
		route:     route,
	}
	config.Set("migrations", migration)
	return s
}

func (s Server) Start() {
	s.autoMigrate()
	s.seed()
	s.startQueue()
	s.startServer()
	s.wait()
}

func (s Server) seed() {
	seeder.Seed(s.db)
}

// autoMigrate 自动迁移
func (s Server) autoMigrate() {
	m, err := migrate.GetInstance()
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}

func (s *Server) startQueue() {
	enable := s.conf.GetBool("queue.enable")
	if enable {
		s.queue = task.NewQueue(50, s.logger)
		s.queue.Start()
		return
	}
	return
}

func (s *Server) startServer() {
	s.httpServer = http.New(s.route, "", s.conf.GetInt("http.port"), s.logger)
	s.httpServer.Start()
	return
}

func (s Server) wait() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	for {
		select {
		case <-c:
			s.httpServer.Stop()
			if s.queue != nil {
				s.queue.Stop()
			}
			return
		}
	}
}
