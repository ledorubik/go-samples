package main

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go-samples/3-http-server-gin/config"
	http_server "go-samples/3-http-server-gin/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Init config
	cfg := config.New()

	// Init logger
	logger := logrus.New()

	// Init HTTP server
	httpServer, err := http_server.NewServer(cfg, logger)

	// Run HTTP server
	go func() {
		if err := httpServer.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("http server start error: %s", err)
		}
	}()

	<-quitChannel
	logger.Info("shutdown os signal")

	// Stop HTTP server
	err = httpServer.Stop()
	if err != nil {
		logger.Errorf("http server shutdown error: %s", err)
	}

}
