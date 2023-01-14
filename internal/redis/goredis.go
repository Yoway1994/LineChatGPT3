package redis

import (
	"fmt"
	"log"
	"sync"

	"context"
	"time"

	"github.com/Yoway1994/LineChatGPT3/config"
	"github.com/Yoway1994/LineChatGPT3/domain"
	"github.com/go-redis/redis/v8"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type GoRedis struct {
	config *config.Config
	conn   *redis.Client
}

var rdb *redis.Client
var rdbLock sync.Mutex

func NewGoRedis(config *config.Config) *GoRedis {
	if rdb == nil {
		rdbLock.Lock()
		defer rdbLock.Unlock()
		// 再次檢查
		if rdb == nil {
			log.Println("connect redis")
			zap.S().Info("connect redis")
			rdb = redis.NewClient(&redis.Options{
				Addr:        fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
				DB:          config.Redis.Database,
				PoolSize:    config.Redis.Max_active,
				IdleTimeout: 1 * time.Minute,
			})
			_, err := rdb.Ping(context.Background()).Result()
			if err != nil {
				log.Println("connect redis error:", errors.WithStack(err))
				zap.S().Info("connect redis error:", errors.WithStack(err))
			} else {
				log.Println("connect redis success")
				zap.S().Info("connect redis success")
			}
		}
	}

	return &GoRedis{
		config: config,
		conn:   rdb,
	}
}

func NewRedis(config *config.Config) domain.GoRedis {
	r := NewGoRedis(config)
	return r
}

func (r *GoRedis) Set(key string, value interface{}) error {
	return r.conn.Set(context.Background(), key, value, 600*time.Second).Err()
}

func (r *GoRedis) Del(key string) error {
	return r.conn.Del(context.Background(), key).Err()
}

func (r *GoRedis) Get(key string) (string, error) {
	return r.conn.Get(context.Background(), key).Result()
}

func (r *GoRedis) Keys(key string) ([]string, error) {
	return r.conn.Keys(context.Background(), key).Result()
}

func (r *GoRedis) INCR(key string) error {
	return r.conn.Incr(context.Background(), key).Err()
}

func (r *GoRedis) INCRBy(key string, num int64) error {
	return r.conn.IncrBy(context.Background(), key, num).Err()
}

func (r *GoRedis) DECR(key string) error {
	return r.conn.Decr(context.Background(), key).Err()
}

func (r *GoRedis) Expire(key string, ttl int) error {
	return r.conn.Expire(context.Background(), key, time.Duration(ttl)*time.Second).Err()
}

func (r *GoRedis) SetEx(key string, value interface{}, ttl int) error {
	return r.conn.SetEX(context.Background(), key, value, time.Duration(ttl)*time.Second).Err()
}

func (r *GoRedis) SetNX(key string, value interface{}) (bool, error) {
	return r.conn.SetNX(context.Background(), key, value, 4*time.Minute).Result()
}

func (r *GoRedis) Lpush(key string, value interface{}) error {
	_, err := redis.NewScript(`
		local key = KEYS[1]
		local change = ARGV[1]
		redis.call("LREM", key, 0, change)
		redis.call("LPUSH", key, change)
	`).Run(context.Background(), r.conn, []string{key}, value).Int()
	if err != nil && err != redis.Nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *GoRedis) LPop(key string) (string, error) {
	return r.conn.LPop(context.Background(), key).Result()
}

func (r *GoRedis) LRange(key string, start, end int64) ([]string, error) {
	return r.conn.LRange(context.Background(), key, start, end).Result()
}

func (r *GoRedis) RPush(key string, value interface{}) error {
	_, err := r.conn.RPush(context.Background(), key, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *GoRedis) RPop(key string) (string, error) {
	return r.conn.RPop(context.Background(), key).Result()
}

func (r *GoRedis) Llen(key string) (int64, error) {
	return r.conn.LLen(context.Background(), key).Result()
}
