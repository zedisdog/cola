package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	svr    *http.Server
	logger *logrus.Logger
}

func New(r *gin.Engine, host string, port int, logger *logrus.Logger) *Server {
	if port == 0 {
		port = 80
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	svr := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	return &Server{
		svr:    svr,
		logger: logger,
	}
}

func (s Server) Start() {
	go func() {
		if err := s.svr.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err.Error())
			}
		}
	}()
}

func (s Server) Stop() {
	t := 10
	cancelCxt, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	defer cancel()
	s.logger.Info(fmt.Sprintf("server will shutdown in %d seconds.", t))
	if err := s.svr.Shutdown(cancelCxt); err != nil {
		s.logger.WithError(err).Error("server shutdown error")
	} else {
		s.logger.Info("http server is shutdown")
	}
}
