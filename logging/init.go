package logging

import (
	"github.com/aka-achu/eidos"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Declaring logger variable for global access
var (
	ControllerLogger *zap.SugaredLogger
	ServiceLogger *zap.SugaredLogger
	RequestLogger    *zap.SugaredLogger
	SQLLogger        *log.Logger
)

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	l, err := eidos.New(fileName, &eidos.Options{
		Size:             10,
		Period:           0,
		RetentionPeriod:  30,
		Compress:         true,
		CompressionLevel: 0,
		LocalTime:        true,
	}, &eidos.Callback{})
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(l)
}

func Initialize() {

	path, _ := os.Getwd()
	path = filepath.Join(path, "log")

	// Validating the existence of the logging path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatalf("[ERROR] Failed to create the folder for storing the logs. Err-%v", err)
		}
	}

	// Initializing the declared logging variables
	ControllerLogger = zap.New(
		zapcore.NewCore(
			getEncoder(),
			getLogWriter(filepath.Join(path, "controller.log")),
			zapcore.InfoLevel,
		),
		zap.AddCaller(),
	).Sugar()
	ServiceLogger = zap.New(
		zapcore.NewCore(
			getEncoder(),
			getLogWriter(filepath.Join(path, "service.log")),
			zapcore.InfoLevel,
		),
		zap.AddCaller(),
	).Sugar()
	RequestLogger = zap.New(
		zapcore.NewCore(
			getEncoder(),
			getLogWriter(filepath.Join(path, "request.log")),
			zapcore.InfoLevel,
		),
		zap.AddCaller(),
	).Sugar()

	l, err := eidos.New(filepath.Join(path, "sql.log"), &eidos.Options{
		Size:             10,
		Period:           24 * time.Hour,
		RetentionPeriod:  7,
		Compress:         true,
		CompressionLevel: 9,
		LocalTime:        true,
	}, &eidos.Callback{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to create log file. Err-%v", err)
	}
	SQLLogger = log.New(l, "", log.LstdFlags|log.Lshortfile)

	log.Println("[INFO] Logging interfaces initialized")
}
