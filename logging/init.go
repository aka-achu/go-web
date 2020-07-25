package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
)

var (
	AppLogger     *zap.SugaredLogger
	RequestLogger *zap.SugaredLogger
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
	})
}
func Initialize() {

	var build = os.Getenv("BUILD")
	var path string
	if build == "Dev" {
		path, _ = os.Getwd()
		path = filepath.Join(path, "log")
	} else if build == "Prod" {
		path = filepath.Join("var", "log", "web")
	} else {
		log.Fatal("Unexpected BUILD value in .env file")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal("Failed to create the folder for storing the logs", err)
		}
	}

	appLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "app.log")), zapcore.InfoLevel)
	requestLoggerCore := zapcore.NewCore(getEncoder(), getLogWriter(filepath.Join(path, "request.log")), zapcore.InfoLevel)

	AppLogger = zap.New(appLoggerCore, zap.AddCaller()).Sugar()
	RequestLogger = zap.New(requestLoggerCore, zap.AddCaller()).Sugar()
}
