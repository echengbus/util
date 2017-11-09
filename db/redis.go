package db

import (
	"fmt"
	"time"

	"driver/conf"

	"github.com/garyburd/redigo/redis"
)

var (
	RedisPool *redis.Pool
	conn      redis.Conn
)

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.REDIS_HOST)
			if err != nil {
				return nil, err
			}

			if conf.REDIS_PASSWORD != "" {
				_, err := c.Do("AUTH", conf.REDIS_PASSWORD)
				if err != nil {
					fmt.Printf("redis password error %v\n", err)
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func RedisGet(key string, database int) string {
	var data []byte

	rc := RedisPool.Get()
	defer rc.Close()
	rc.Do("SELECT", database)

	data, err := redis.Bytes(rc.Do("GET", key))
	if err != nil {
		return ""
	}
	return string(data)
}

func RedisSet(key string, value string, database int) {
	rc := RedisPool.Get()
	defer rc.Close()
	rc.Do("SELECT", database)
	rc.Do("SET", key, value)
}

func RedisLock(tag string, sec int) {
	rc := RedisPool.Get()
	defer rc.Close()
	rc.Do("SELECT", 0)
	keyName := "redis-lock:" + tag
	rc.Do("SET", keyName, 1)
	rc.Do("EXPIRE", keyName, sec)
}

func RedisLockCheck(tag string) bool {
	rc := RedisPool.Get()
	defer rc.Close()
	rc.Do("SELECT", 0)

	keyName := "redis-lock:" + tag
	_, err := redis.Bytes(rc.Do("GET", keyName))
	if err != nil {
		return false
	}
	return true
}
