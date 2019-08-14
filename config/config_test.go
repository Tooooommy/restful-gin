package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := Get()
	fmt.Println(cfg.App.Mode)
	if cfg == nil {
		t.Errorf("test get error")
	}
}
