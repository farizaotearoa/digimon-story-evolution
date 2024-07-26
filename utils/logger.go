package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger() {
	fmt.Println("Logger Path (", ConfigRootKey+LoggerPath, ") :", Config.GetString(LoggerPath))
	logDir := fmt.Sprintf("%s/%s", Config.GetString(LoggerPath), Config.GetString(AppsName))
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	timestamp := time.Now().Format("2006-01-02")
	logFile := fmt.Sprintf("%s/%s.log", logDir, timestamp)

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.FullCallerEncoder,
	}

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(file),
		zapcore.InfoLevel,
	)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	core := zapcore.NewTee(fileCore, consoleCore)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.FatalLevel))
	defer Logger.Sync()
}
