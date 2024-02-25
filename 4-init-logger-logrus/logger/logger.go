package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const defaultLogLevel = logrus.DebugLevel

// InitLogger initializes logger
// if filename is empty, logger's output sets to stdout
func InitLogger(level, filename string, enableReportCaller, hidePrivateInfo bool, privateWords ...string) (*logrus.Logger, error) {
	logger := logrus.New()

	if len(filename) != 0 {
		logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return nil, fmt.Errorf("logfile opening error: %v", err.Error())
		}
		logger.SetOutput(logFile)
	} else {
		logger.SetOutput(os.Stdout)
	}

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logger.SetFormatter(formatter)
	logger.SetReportCaller(enableReportCaller)

	l, err := logLevelConvertor(level)
	if err != nil {
		return nil, err
	}

	logger.Infof("log level: %s", l.String())

	logger.SetLevel(l)

	if hidePrivateInfo {
		logger.AddHook(&hidePrivateInfoHook{privateWords: privateWords})
	}

	return logger, nil
}

func logLevelConvertor(l string) (logrus.Level, error) {
	l = strings.ToLower(l)
	switch l {
	case "":
		return defaultLogLevel, nil
	case "trace":
		return logrus.TraceLevel, nil
	case "debug":
		return logrus.DebugLevel, nil
	case "info":
		return logrus.InfoLevel, nil
	case "warning":
		return logrus.WarnLevel, nil
	case "error":
		return logrus.ErrorLevel, nil
	case "fatal":
		return logrus.FatalLevel, nil
	case "panic":
		return logrus.PanicLevel, nil
	default:
		return 0, fmt.Errorf("incorrect log level")
	}
}
