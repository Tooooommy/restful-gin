package logger

import (
	"github.com/natefinch/lumberjack"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger        *zap.Logger
	Sugar         *zap.SugaredLogger
	defaultOutput = "log/im.log"
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

func SetOutput(path string) {
	defaultOutput = path
}

func init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   defaultOutput,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			w),
		zap.DebugLevel,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}
