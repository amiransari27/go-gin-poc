package logger

import (
	"fmt"
	"go-gin-api/src/config"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logrus interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
}

type logrus struct {
	logger map[string]*log.Logger
}

var logLevels = map[string]log.Level{
	"Info":  log.InfoLevel,
	"Debug": log.DebugLevel,
	"Warn":  log.WarnLevel,
	"Error": log.ErrorLevel,
}

func NewLogrus() Logrus {
	logDir := config.GetConfig().LogDir
	logInstance := &logrus{logger: make(map[string]*log.Logger)}

	for severity, logLevel := range logLevels {
		lg := log.New()
		lg.SetFormatter(&log.JSONFormatter{})
		lg.SetLevel(logLevel)
		file, err := os.OpenFile(
			fmt.Sprintf("%s/%s.log", logDir, severity), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666,
		)
		if err == nil {
			lg.SetOutput(file)
		} else {
			log.Fatal("Error occured while creating logger file", err)
		}
		logInstance.logger[severity] = lg
	}

	return logInstance
}

func (l *logrus) Debug(args ...interface{}) {
	l.logger["Debug"].Debug(args...)
}

func (l *logrus) Info(args ...interface{}) {
	l.logger["Info"].Info(args...)
}
func (l *logrus) Warning(args ...interface{}) {
	l.logger["Warn"].Warn(args...)
}
func (l *logrus) Error(args ...interface{}) {
	l.logger["Error"].Error(args...)
}
