package logger

import (
	"CrownDaisy_GOGIN/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger *logrus.Logger

// info
func Info(args ...interface{}) {
	if logger != nil {
		logger.Info(args...)
	}
}

func Infof(format string, args ...interface{}) {
	if logger != nil {
		logger.Infof(format, args...)
	}
}

func Infoln(args ...interface{}) {
	if logger != nil {
		logger.Infoln(args...)
	}
}

// debug
func Debug(args ...interface{}) {
	if logger != nil {
		logger.Debug(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if logger != nil {
		logger.Debugf(format, args...)
	}
}

func Debugln(args ...interface{}) {
	if logger != nil {
		logger.Debugln(args...)
	}
}

// warn
func Warn(args ...interface{}) {
	if logger != nil {
		logger.Warn(args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if logger != nil {
		logger.Warnf(format, args...)
	}
}

func Warnln(args ...interface{}) {
	if logger != nil {
		logger.Warnln(args...)
	}
}

// warning
func Warning(args ...interface{}) {
	if logger != nil {
		logger.Warning(args...)
	}
}
func Warningf(format string, args ...interface{}) {
	if logger != nil {
		logger.Warningf(format, args...)
	}
}

func Warningln(args ...interface{}) {
	if logger != nil {
		logger.Warningln(args...)
	}
}

// Error
func Error(args ...interface{}) {
	if logger != nil {
		logger.Error(args...)
	}
}
func Errorf(format string, args ...interface{}) {
	if logger != nil {
		logger.Errorf(format, args...)
	}
}

func Errorln(args ...interface{}) {
	if logger != nil {
		logger.Errorln(args...)
	}
}

// trace
func Trace(args ...interface{}) {
	if logger != nil {
		logger.Trace(args...)
	}
}
func Tracef(format string, args ...interface{}) {
	if logger != nil {
		logger.Tracef(format, args...)
	}
}

func Traceln(args ...interface{}) {
	if logger != nil {
		logger.Traceln(args...)
	}
}

// Fatal
func Fatal(args ...interface{}) {
	if logger != nil {
		logger.Fatal(args...)
	}
}
func Fatalf(format string, args ...interface{}) {
	if logger != nil {
		logger.Fatalf(format, args...)
	}
}

func Fatalln(args ...interface{}) {
	if logger != nil {
		logger.Fatalln(args...)
	}
}

// panic

func Panic(args ...interface{}) {
	if logger != nil {
		logger.Panic(args...)
	}
}
func Panicf(format string, args ...interface{}) {
	if logger != nil {
		logger.Panicf(format, args...)
	}
}

func Panicln(args ...interface{}) {
	if logger != nil {
		logger.Panicln(args...)
	}
}

// print
func Print(args ...interface{}) {
	if logger != nil {
		logger.Print(args...)
	}
}
func Printf(format string, args ...interface{}) {
	if logger != nil {
		logger.Printf(format, args...)
	}
}

func Println(args ...interface{}) {
	if logger != nil {
		logger.Warningln(args...)
	}
}

func init() {
	InitLogger()
}

func InitLogger() {
	logger = logrus.New()
	cfg := config.Get().Logger
	logger.SetLevel(getLevelByConf(cfg.Level))
	if cfg.Formatter == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				"FieldKeyTime":  "@timestamp",
				"FieldKeyLevel": "@level",
				"FieldKeyMsg":   "@message",
			},
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
	logger.SetOutput(getOutputByConf(cfg.Output))
	logger.SetReportCaller(true)
}

func getLevelByConf(level string) logrus.Level {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return logrus.InfoLevel
	}
	return l
}

func getOutputByConf(out string) io.Writer {
	switch out {
	case "stdin":
		return os.Stdin
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	case "":
		out = "app.log"
		fallthrough
	default:
		file, err := os.Open(out)
		if err != nil {
			return nil
		}
		return file
	}
}
