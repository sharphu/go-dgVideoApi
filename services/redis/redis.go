package redisClient

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

//直连
func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", beego.AppConfig.String("redisdb"))

	return pool
}

//通过连接池链接
func PoolConnect() redis.Conn {
	pool := &redis.Pool{
		MaxIdle: 5000, //最大空闲连接数
		MaxActive: 10000, //最大连接数
		IdleTimeout: 180*time.Second, //空闲连接超时时间
		Wait: true, //超过最大连接数时，保持等待
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	return pool.Get()
}
