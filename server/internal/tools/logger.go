package tools

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(file *os.File) (*zap.Logger, error) {
	cfg := zap.Config{
		Encoding:         "console", // Используем консольный кодировщик
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{file.Name()},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
			TimeKey:    "time",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("[2006-01-02 15:04:05 [MSK]")) // Формат времени
			},
			EncodeLevel: zapcore.CapitalLevelEncoder, // Формат для уровня
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("сan't init Zap logger: %w", err)
	}
	return logger, nil
}

func createDirLog(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create dirlog: %w", err)
	}
	return nil
}

func InitLogFile() (*os.File, error) {
	now := time.Now().Format("2006-01-02")
	path := os.Getenv("LOG_PATH")
	fileName := path + now + ".log"

	if err := createDirLog(path); err != nil {
		return nil, err
	}

	var file *os.File
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err = os.Create(fileName)
		if err != nil {
			return nil, fmt.Errorf("сan't create a file: %w", err)
		}
	} else {
		file, err = os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("can't open a file: %w", err)
		}
	}
	return file, nil
}
