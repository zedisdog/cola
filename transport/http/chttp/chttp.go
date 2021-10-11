package chttp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var httpServers = make(map[string]*Svr)

func Server(name ...string) *Svr {
	serverName := "default"
	if len(name) > 0 {
		serverName = name[0]
	}

	return getOrCreateServer(serverName)
}

func getOrCreateServer(name string) *Svr {
	if server, ok := httpServers[name]; ok {
		return server
	}

	httpServers[name] = New()
	return httpServers[name]
}

func New() *Svr {
	return &Svr{
		svr: &http.Server{},
	}
}

type Svr struct {
	svr    *http.Server
	logger *logrus.Logger
}

func (s *Svr) SetRouter(r *gin.Engine) {
	s.svr.Handler = r
}

func (s *Svr) SetAddr(addr string) {
	s.svr.Addr = addr
}

func (s *Svr) SetLogger(logger *logrus.Logger) {
	s.logger = logger
}

func (s *Svr) Start() {
	go func() {
		if err := s.svr.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err.Error())
			}
		}
	}()
}

func (s *Svr) Stop() {
	t := 10
	cancelCxt, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	defer cancel()

	s.logInfo(fmt.Sprintf("server will shutdown in %d seconds.", t))
	if err := s.svr.Shutdown(cancelCxt); err != nil {
		s.logError(err, "server shutdown error")
	} else {
		s.logInfo("http server is shutdown")
	}
}

func (s Svr) logInfo(msg string) {
	if s.logger != nil {
		s.logger.Info(msg)
	}
}

func (s Svr) logError(err error, msg string) {
	if s.logger != nil {
		s.logger.WithError(err).Error(msg)
	}
}
