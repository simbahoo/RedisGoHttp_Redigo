package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	//"strconv"
)

var (
	// 定义常量
	RedisClient    *redis.Pool
	REDIS_HOST     string
	REDIS_DB       int
	REDIS_PASSWORD string
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = beego.AppConfig.String("redis.host")
	REDIS_DB, _ = beego.AppConfig.Int("redis.db")
	REDIS_PASSWORD = beego.AppConfig.String("redis.password")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 1),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 10),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}

			if REDIS_PASSWORD != "" {
				if _, err := c.Do("AUTH", REDIS_PASSWORD); err != nil {

					return nil, err
				}
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

var (
	RedisList map[string]*Redis
)

type Redis struct {
	Cmd   string
	Key   string
	Value string
}

func SetRedis(r Redis) (o Redis, err error) {
	// 从池里获取连接
	rc := RedisClient.Get()

	// 用完后将连接放回连接池
	defer rc.Close()
	switch strings.ToUpper(r.Cmd) {

	case "SET":
		reply, err := rc.Do("SET", r.Key, r.Value)
		fmt.Println(reply)
		return r, err
	case "GET":
		reply, err := redis.String(rc.Do("GET", r.Key))
		fmt.Println(reply)
		r.Value = reply
		return r, err
	default:
		return r, errors.New("Cmd UnKnown")
	}
}

func GetRedis(uid string) (u *Redis, err error) {

	if u, ok := RedisList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Redis not exists")
}

func GetAllRediss() map[string]*Redis {
	return RedisList
}
