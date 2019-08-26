package logger

import (
	"github.com/natefinch/lumberjack"
	"os"
	"restful-gin/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func InitLogger() {
	cfg := config.Get().Logger
	if cfg.Output == "" {
		cfg.Output = "log/im.log"
	}
	if cfg.MaxAge == 0 {
		cfg.MaxAge = 28
	}
	if cfg.MaxSize == 0 {
		cfg.MaxSize = 500
	}
	if cfg.MaxBackups == 0 {
		cfg.MaxBackups = 3
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.Output,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.DebugLevel,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}

func init() {
	InitLogger()
}
