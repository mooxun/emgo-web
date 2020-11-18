package redis

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/logger"
)

var pool *redis.Pool

func Init() {
	dialFunc := func() (c redis.Conn, err error) {
		dsn := fmt.Sprintf("%s:%d", cfg.App.Redis.Host, cfg.App.Redis.Port)
		c, err = redis.Dial("tcp", dsn)

		if err != nil {
			logger.Error("redis connect error.")
			panic(err.Error())
		}
		password := cfg.App.Redis.Password
		if password != "" {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				logger.Error("redis invalid password error.")
				panic(err.Error())
			}
		}
		return
	}

	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial:        dialFunc,
	}

	logger.Info("redis connect success")
}

// actually do the redis cmds, args[0] must be the key name.
func do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	c := pool.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// Put put value to redis.
func Put(key string, val interface{}, timeout int) error {
	_, err := do("SETEX", key, timeout, val)
	return err
}

// Get value from redis.
func Get(key string) string {
	if v, err := redis.String(do("GET", key)); err == nil {
		return v
	}
	return ""
}

// Delete delete value in redis.
func Delete(key string) error {
	_, err := do("DEL", key)
	return err
}

// IsExist check cache's existence in redis.
func IsExist(key string) bool {
	v, err := redis.Bool(do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

// Incr increase counter in redis.
func Incr(key string) error {
	_, err := redis.Bool(do("INCRBY", key, 1))
	return err
}

// Decr decrease counter in redis.
func Decr(key string) error {
	_, err := redis.Bool(do("INCRBY", key, -1))
	return err
}

// lpush value to redis.
func Lpush(key string, val string) error {
	_, err := do("LPUSH", key, val)
	return err
}

// rpop value from redis.
func Rpop(key string) string {
	v, err := redis.String(do("RPOP", key))
	if err == nil {
		return v
	}
	return ""
}
