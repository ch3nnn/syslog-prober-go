package common

import "github.com/gomodule/redigo/redis"

var cachePool *redis.Pool

func CacheInitialization() {

	// 实例化一个连接池
	cachePool = &redis.Pool{
		MaxIdle:     16,  // 最初连接数量
		MaxActive:   0,   // 连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, // 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { // 要连接的redis数据库
			return redis.Dial(
				Configuration().CacheNetwork,
				Configuration().CacheAddress,
				redis.DialPassword(Configuration().CachePassword),
				redis.DialDatabase(Configuration().CacheDatabase),
			)
		},
	}
}

// CachePool 获取连接池
func CachePool() *redis.Pool {
	return cachePool
}
