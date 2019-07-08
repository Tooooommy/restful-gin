package config

import (
	"github.com/go-ini/ini"
)

var DefaultConfigPath = "app.ini"

type Mysql struct {
	Host              string `ini:"host"`
	Username          string `ini:"username"`
	Password          string `ini:"password"`
	Schema            string `ini:"schema"`
	Charset           string `ini:"charset"`
	Loc               string `ini:"loc"`
	MaxIdleConns      int    `ini:"max_idle_conns"`
	MaxOpenConns      int    `ini:"max_open_conns"`
	MaxConnLifetime   int    `ini:"max_conn_lifetime"`
	ConnectionTimeout int    `ini:"connection_timeout"`
}

type Redis struct {
	Host        string `ini:"host"`
	Password    string `ini:"password"`
	Db          int    `ini:"db"`
	MaxIdle     int    `ini:"max_idle"`
	MaxActive   int    `ini:"max_active"`
	IdleTimeout int    `ini:"idle_timeout"`
	Wait        bool   `ini:"wait"`
}

type Logger struct {
	Output    string `ini:"output"`
	Formatter string `ini:"formatter"`
	Level     string `ini:"level"`
}

type Jwt struct {
	Issuer  string `ini:"issuer"`
	Secret  string `ini:"secret"`
	Expired int    `ini:"expired"`
}

type Pool struct {
	Ants   int `ini:"ants"`
	Worker int `ini:"worker"`
	Job    int `ini:"job"`
}

type Spider struct {
	StartUrl []string `ini:"url"`
}

type WeChat struct {
	AppId       string `ini:"app_id"`
	AppSecret   string `ini:"app_secret"`
	RedirectUri string `ini:"redirect_uri"`
	Scope       string `ini:"scope"`
	State       string `ini:"state"`
	Lang        string `ini:"lang"`
}

type QQ struct {
	ClientId     string `ini:"client_id"`
	ClientSecret string `ini:"client_secret"`
	RedirectUri  string `ini:"redirect_uri"`
	Scope        string `ini:"scope"`
	Display      string `ini:"display"`
	State        string `ini:"state"`
}

type Config struct {
	AppMode   string `ini:"app_mode"`
	PwdSecret string `ini:"pwd_secret"`
	Mysql     `ini:"mysql"`
	Redis     `ini:"redis"`
	Logger    `ini:"logger"`
	Jwt       `ini:"jwt"`
	Pool      `ini:"pool"`
	WeChat    `ini:"we_chat"`
	QQ        `ini:"qq"`
}

var cfg *Config

func InitConf() error {
	cfg = new(Config)
	return ini.MapTo(&cfg, DefaultConfigPath)
}

func Get() *Config {
	if cfg == nil {
		if err := InitConf(); err != nil {
			panic(err)
		}
	}
	return cfg
}
