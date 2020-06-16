package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

type Mysql struct {
	Host              string `json:"host"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	Schema            string `json:"schema"`
	Charset           string `json:"charset"`
	Loc               string `json:"loc"`
	MaxIdleConn     int    `json:"max_idle_conns"`
	MaxOpenConn      int    `json:"max_open_conns"`
	MaxConnLifetime   int    `json:"max_conn_lifetime"`
	ConnectionTimeout int    `json:"connection_timeout"`
}

type Redis struct {
	Host        string `json:"host"`
	Password    string `json:"password"`
	Db          int    `json:"db"`
	MaxIdle     int    `json:"max_idle"`
	MaxActive   int    `json:"max_active"`
	IdleTimeout int    `json:"idle_timeout"`
	Wait        bool   `json:"wait"`
}

type Logger struct {
	Output     string `json:"output" default:"a"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Level      string `json:"level"`
}

type Jwt struct {
	Subject   string `json:"subject"`
	Audience  string `json:"audience"`
	NotBefore int64  `json:"not_before"`
	Issuer    string `json:"issuer"`
	Secret    string `json:"secret"`
	Expired   int    `json:"expired"`
}

type Pool struct {
	Ants   int `json:"ants"`
	Worker int `json:"worker"`
	Job    int `json:"job"`
}

type Spider struct {
	StartUrl []string `json:"url"`
}

type WeChat struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	RedirectUri string `json:"redirect_uri"`
	Scope       string `json:"scope"`
	State       string `json:"state"`
	Lang        string `json:"lang"`
}

type QQ struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectUri  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	Display      string `json:"display"`
	State        string `json:"state"`
}
type App struct {
	Mode   string `json:"mode"`
	Port   string `json:"port"`
	Secret string `json:"secret"`
}
type Config struct {
	App    `json:"app"`
	Mysql  `json:"mysql"`
	Redis  `json:"redis"`
	Logger `json:"logger"`
	Jwt    `json:"jwt"`
	Pool   `json:"pool"`
	WeChat `json:"we_chat"`
	QQ     `json:"qq"`
}

var (
	cfg               *Config
	configFile = "config.toml"
)

func init()  {
	err := InitConfig()
	if err != nil {
		panic(fmt.Sprintf("init config: %v", err))
	}
}

func Get() *Config {
	if cfg == nil {
		_ = InitConfig()
	}
	return cfg
}
func InitConfig(path ...string) error {
	if len(path) > 0 {
		configFile = path[0]
	}
	cfg = &Config{}

	viper.SetConfigType(filepath.Ext(configFile)[1:])
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		_, file, _, _ := runtime.Caller(1)
		fmt.Print(filepath.Dir(file))
		configFile = filepath.Join(filepath.Dir(file), configFile)
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}
	return nil
}
