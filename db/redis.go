package db

import (
	"CrownDaisy_GOGIN/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

type RedisDB struct {
	Host        string
	Password    string
	Db          int
	Pool        *redis.Pool
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	Wait        bool
}

func (r *RedisDB) init() error {
	r.Pool = &redis.Pool{
		MaxIdle:   r.MaxIdle,
		MaxActive: r.MaxActive,
		Wait:      r.Wait,
		Dial: func() (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", r.Host)
			if err != nil {
				fmt.Println(err)
				return
			}
			if r.Password != "" {
				if _, err := conn.Do("AUTH", r.Password); err != nil {
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
	return nil
}

// == pool
func (r *RedisDB) GetConn() redis.Conn {
	if r.Pool == nil {
		if err := r.init(); err != nil {
			return nil
		}
	}
	conn := r.Pool.Get()
	_, _ = conn.Do("SELECT", r.Db)
	return conn
}

func (r *RedisDB) PoolActiveCount() int {
	return r.Pool.ActiveCount()
}

func (r *RedisDB) PoolClose() error {
	return r.Pool.Close()
}

func (r *RedisDB) PoolIdleCount() int {
	return r.Pool.IdleCount()
}

// == cmd
func (r *RedisDB) Ping() (string, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.String(conn.Do("PING"))
}

// == string
func (r *RedisDB) Set(key string, val string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SET", key, val)
	return err
}

func (r *RedisDB) Get(key string) (string, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func (r *RedisDB) GetInt(key string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

func (r *RedisDB) Unshift(key string, val string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("LPUSH", key, val)
	return err
}

func (r *RedisDB) Setex(key string, val string, time int) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SETEX", key, time, val)
	return err
}

func (r *RedisDB) SetInt(key string, val int) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SET", key, val)
	return err
}

func (r *RedisDB) Del(key string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

func (r *RedisDB) Incr(key string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

func (r *RedisDB) IncrBy(key string, step int) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("INCRBY", key, step))
}

func (r *RedisDB) Decr(key string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("DECR", key))
}

func (r *RedisDB) DecrBy(key string, step int) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("DECRBY", key, step))
}

func (r *RedisDB) Exists(key string) (bool, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

func (r *RedisDB) Expire(key string, expire int) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, expire)
	return err
}

func (r *RedisDB) GetTTL(key string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("ttl", key))
}

//Hashes
func (r *RedisDB) HSET(key, field string, value interface{}) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("HSET", key, field, value)
	return err
}

func (r *RedisDB) HLen(key string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("HLEN", key))
}

func (r *RedisDB) HExists(key, field string) (bool, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Bool(conn.Do("HEXISTS", key, field))
}

func (r *RedisDB) HGet(key, field string) (string, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.String(conn.Do("HGET", key, field))
}

func (r *RedisDB) HGetAll(key string) ([]string, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Strings(conn.Do("HGETALL", key))
}

func (r *RedisDB) HGetInt(key, field string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("HGET", key, field))
}

func (r *RedisDB) HDel(key string, field string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("HDEL", key, field)
	return err
}

// set
func (r *RedisDB) SAdd(key, member string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SADD", key, member)
	return err
}

func (r *RedisDB) SAddInt(key string, member int) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SADD", key, member)
	return err
}

func (r *RedisDB) SCard(key string) (int, error) {
	conn := r.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("SCARD", key))
}

func (r *RedisDB) SRemInt(key string, member int) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SREM", key, member)
	return err
}

func (r *RedisDB) SRem(key, member string) error {
	conn := r.GetConn()
	defer conn.Close()
	_, err := conn.Do("SREM", key, member)
	return err
}

func (r *RedisDB) SMembers(key string) ([]string, error) {
	var resArr = make([]string, 0)
	conn := r.GetConn()
	defer conn.Close()
	bArr, err := redis.ByteSlices(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	for k := range bArr {
		resArr = append(resArr, string(bArr[k]))
	}
	return resArr, err
}

func (r *RedisDB) SMembersInt(key string) ([]int, error) {
	var resArr = make([]int, 0)
	conn := r.GetConn()
	defer conn.Close()
	bArr, err := redis.ByteSlices(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}

	for k := range bArr {
		var val int
		val, _ = strconv.Atoi(string(bArr[k]))
		resArr = append(resArr, val)
	}
	return resArr, err
}

// list
func (r *RedisDB) LPush(key string, value string) error {
	conn := r.GetConn()
	defer conn.Close()

	_, err := conn.Do("LPUSH", key, value)
	return err
}

func (r *RedisDB) RPop(key string) (string, error) {
	conn := r.GetConn()
	defer conn.Close()

	return redis.String(conn.Do("RPOP", key))
}

func (r *RedisDB) BRPop(key string, timeout int) (string, error) {
	conn := r.GetConn()
	defer conn.Close()
	res, err := redis.ByteSlices(conn.Do("BRPOP", key, timeout))
	if err != nil {
		return "", err
	}
	if len(res) < 2 {
		return "", err
	}
	return string(res[1]), nil
}

// PubSub
func (r *RedisDB) Publish(channel string, msg string) error {
	conn := r.GetConn()
	defer conn.Close()
	err := conn.Send("PUBLISH", channel, msg)
	if err != nil {
		return err
	}
	return conn.Flush()
}

func (r *RedisDB) GetSubConn() redis.PubSubConn {
	conn := redis.PubSubConn{
		Conn: r.GetConn(),
	}
	return conn
}

var RDB *RedisDB

func InitRedisDB() error {
	cfg := config.Get().Redis
	RDB = &RedisDB{
		Host:        cfg.Host,
		Password:    cfg.Password,
		Db:          cfg.Db,
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout,
		Wait:        cfg.Wait,
	}
	err := RDB.init()
	return err
}

func GetRedisDB() *RedisDB {
	return RDB
}
