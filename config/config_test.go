package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	//var path string
	//pwd, _ := os.Getwd()
	//pwdParts := strings.Split(pwd, "/")
	//if len(pwdParts) > 2 {
	//	path = strings.Join(pwdParts[:len(pwdParts)-1], "/")
	//}
	//path = path + "/app.ini"

	DefaultConfigPath = "app.ini"
	// == test InitConf

	// == test Get
	cfg := Get()
	if cfg == nil || cfg.AppMode != "dev" {
		t.Errorf("test get error")
	}

	// test cfg mysql
	fmt.Println(cfg.Mysql)
}
