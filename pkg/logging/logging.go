package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type Logger struct {
	InfoLogger  *logrus.Logger
	ErrorLogger *logrus.Logger
}

func setConfiguration(logger *logrus.Logger) {
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02.01.2006 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			dir, fileName := filepath.Split(f.File)
			dir = strings.TrimSuffix(dir, string(os.PathSeparator)) // reduce possible slash in the end of dirname
			lastDir := path.Base(dir)
			file = fmt.Sprintf("%s/%s:%d", lastDir, fileName, f.Line)
			return f.Function, file
		},
	})

}

func InitLogger() *Logger {
	infoLogger := logrus.New()
	errorLogger := logrus.New()
	setConfiguration(infoLogger)
	setConfiguration(errorLogger)
	infoLoggerFile, err := os.OpenFile("cmd/logs/logs.info", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		panic("unable to open info file for logs")
	}
	errorLoggerFile, err := os.OpenFile("cmd/logs/logs.error", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil {
		panic("unable to open error file for logs")
	}
	infoLogger.Out = infoLoggerFile
	errorLogger.Out = errorLoggerFile
	return &Logger{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
	}
}