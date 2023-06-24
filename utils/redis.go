package utils

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	// 定义常量
	veriCodePool *redis.Pool
	cookiePool   *redis.Pool
	REDIS_HOST   = "127.0.0.1:6379"
	VeriCODE_DB  = 0
	COOKIE_DB    = 1
)

func init() {
	// 建立连接池
	veriCodePool = &redis.Pool{
		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool.
		//最大活跃连接数，0代表无限
		MaxActive: 888,
		//最大闲置连接数
		// Maximum number of idle connections in the pool.
		MaxIdle: 20,
		//闲置连接的超时时间
		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: time.Second * 100,
		//定义拨号获得连接的函数
		// Dial is an application supplied function for creating and configuring a
		// connection.
		//
		// The connection returned from Dial must not be in a special state
		// (subscribed to pubsub channel, transaction started, ...).
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", REDIS_HOST)
		},
	}

	cookiePool = &redis.Pool{
		MaxActive:   888,
		MaxIdle:     20,
		IdleTimeout: time.Second * 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", COOKIE_DB)
			return c, nil
		},
	}

	// //延迟关闭连接池
	// defer pool.Close()

}
