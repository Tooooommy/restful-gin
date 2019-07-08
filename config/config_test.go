package config

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	var path string
	pwd, _ := os.Getwd()
	pwdParts := strings.Split(pwd, "/")
	if len(pwdParts) > 2 {
		path = strings.Join(pwdParts[:len(pwdParts)-1], "/")
	}
	path = path + "/app.ini"

	DefaultConfigPath = path
	// == test InitConf
	err := InitConf()
	if err != nil {
		t.Errorf("test init conf error: %v", err)
	}

	// == test Get
	cfg := Get()
	if cfg == nil || cfg.AppMode != "dev" {
		t.Errorf("test get error: %v", err)
	}

	// test cfg mysql
	fmt.Println(cfg.Mysql)
}
