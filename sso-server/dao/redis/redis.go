/**
 * @Author：Robby
 * @Date：2022/2/4 22:48
 * @Function：
 **/

package redis
import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sso/settings"
)


var Client *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {

	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})

	_, err = Client.Ping().Result()
	if err != nil {
		log.Fatalf("Redis ping failed: %v\n", err)
	}

	return nil
}

func Close() {
	_ = Client.Close()
}