package main

import (
	_logger "go-samples/4-init-logger-logrus/logger"
	"log"
	"strings"
)

const (
	logLevel              = "debug"
	logFile               = ""
	logEnableReportCaller = false
	logHidePrivateInfo    = false
	logPrivateWords       = "password,token,secret"
)

func main() {
	// Init logger
	privateWords := strings.Split(logPrivateWords, ",")

	logger, err := _logger.InitLogger(logLevel, logFile, logEnableReportCaller, logHidePrivateInfo, privateWords...)
	if err != nil {
		log.Fatalf("logger init error: %s", err)
	}

	logger.Info("logger initialized")
}
