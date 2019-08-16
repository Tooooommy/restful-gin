package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Init()
	Logger.Info("hello")
	Sugar.Info("world")
}
