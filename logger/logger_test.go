package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	InitLogger()
	Logger.Info("hello")
	Sugar.Info("world")
}
