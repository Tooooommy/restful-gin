package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Logger.Info("hello")
	Sugar.Info("world")
}
