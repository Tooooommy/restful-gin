package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"restful-gin/config"
	"testing"
	"time"
)

func TestConnRedis(t *testing.T) {
	cfg := config.Get().Redis
	pool := &redis.Pool{
		MaxIdle:   cfg.MaxIdle,
		MaxActive: cfg.MaxActive,
		Wait:      cfg.Wait,
		Dial: func() (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", cfg.Host)
			if err != nil {
				fmt.Println(err)
				return
			}
			if cfg.Password != "" {
				if _, err := conn.Do("AUTH", cfg.Password); err != nil {
					_ = conn.Close()
					return
				}
			}
			return
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		},
	}
	defer pool.Close()

	fmt.Println(cfg.Password)
	conn := pool.Get()
	if conn == nil {
		t.Errorf("get connect from redis pool error: %v", conn)
	}
	defer conn.Close()
	reply, err := conn.Do("PING")
	if err != nil {
		t.Errorf("redis conn error: %+v", err)
	}
	fmt.Println("ping return", reply)

}

func TestInitPool(t *testing.T) {
	err := InitRedisDB()
	if err != nil {
		t.Errorf("init redis db pool error:%v", err)
	}
	r := GetRedisDB()
	conn := r.GetConn()
	if conn == nil {
		t.Errorf("get init db connect error: %v", err)
	}

}

func TestDoRedisCmd(t *testing.T) {

}
