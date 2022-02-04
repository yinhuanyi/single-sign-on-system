/**
 * @Author：Robby
 * @Date：2022/2/4 22:48
 * @Function：
 **/

package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"goods-server/settings"
	"log"
)


var client *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {

	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatalf("Redis连接失败: %v\n", err)
	}

	return nil
}

func Close() {
	_ = client.Close()
}