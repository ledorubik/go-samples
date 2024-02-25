package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-samples/3-http-server-gin/config"
	"net/http"
	"runtime/debug"
	"time"
)

type Server struct {
	server      *http.Server
	loggerEntry *logrus.Entry
}

func NewServer(cfg *config.Config, logger *logrus.Logger) (*Server, error) {
	loggerEntry := logger.WithField("service_name", "http_server")

	if cfg.HttpServerGinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(loggingMiddleware(logger))
	router.Use(httpRecover(loggerEntry))

	RegisterHttpEndpoints(logger, router)

	tlsConfig, err := tlsConfig(cfg.HttpServerTlsEnable, cfg.HttpServerTlsServerCertPath, cfg.HttpServerTlsServerKeyPath)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:      fmt.Sprintf("%s:%d", cfg.HttpServerRestHost, cfg.HttpServerRestPort),
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	return &Server{
		server:      httpServer,
		loggerEntry: loggerEntry,
	}, nil
}

func (s *Server) Start() error {
	return s.run()
}

func (s *Server) Stop() error {
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("http server shutdown error: %s", err.Error())
	}

	s.loggerEntry.Info("http server gracefully shutdown")

	return nil
}

func (s *Server) run() error {
	if s.server.TLSConfig == nil {
		return s.server.ListenAndServe()
	}
	return s.server.ListenAndServeTLS("", "")
}

func tlsConfig(tlsEnable bool, certPath, keyPath string) (*tls.Config, error) {
	if !tlsEnable {
		return nil, nil
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	tc := &tls.Config{}
	tc.Certificates = []tls.Certificate{cert}

	return tc, nil
}

func httpRecover(loggerEntry *logrus.Entry) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				loggerEntry.Warningf("%s panic err %s", ctx.Request.URL, err)
				loggerEntry.Warningf("----------panic stack start----------")
				loggerEntry.Warningf("%s", debug.Stack())
				loggerEntry.Warningf("----------panic stack end----------")

				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
