package logger

import (
	"fmt"
	"go-gin-api/src/appConst"
	"go-gin-api/src/config"
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ILogrus interface {
	Debug(*gin.Context, ...interface{})
	Info(*gin.Context, ...interface{})
	Warning(*gin.Context, ...interface{})
	Error(*gin.Context, ...interface{})
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

func NewLogrus() ILogrus {
	env := strings.ToUpper(os.Getenv("env"))
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
			if env == "PROD" || env == "STAGE" {
				lg.SetOutput(file)
			} else {
				lg.SetOutput(io.MultiWriter(file, os.Stdout))
			}

		} else {
			log.Fatal("Error occured while creating logger file", err)
		}
		logInstance.logger[severity] = lg
	}

	return logInstance
}

func (l *logrus) Debug(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		l.logger["Debug"].WithFields(getLogFields(ctx)).Info(args...)
	} else {
		l.logger["Debug"].Info(args...)
	}
}

func (l *logrus) Info(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		l.logger["Info"].WithFields(getLogFields(ctx)).Info(args...)
	} else {
		l.logger["Info"].Info(args...)
	}

}
func (l *logrus) Warning(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		l.logger["Warn"].WithFields(getLogFields(ctx)).Warn(args...)
	} else {
		l.logger["Warn"].Warn(args...)
	}

}
func (l *logrus) Error(ctx *gin.Context, args ...interface{}) {
	if ctx != nil {
		l.logger["Error"].WithFields(getLogFields(ctx)).Warn(args...)
	} else {
		l.logger["Error"].Warn(args...)
	}
}

func getLogFields(ctx *gin.Context) log.Fields {
	requestId, _ := ctx.Get(appConst.XRequestId)
	fields := log.Fields{
		"request": requestId.(string),
	}
	return fields
}
