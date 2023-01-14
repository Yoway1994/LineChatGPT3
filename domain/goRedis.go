//go:generate ~/go/bin/mockgen -source ./goRedis.go -destination ../mock/redis_mock.go -package domain_mock

package domain

type GoRedis interface {
	Set(key string, value interface{}) error
	Del(key string) error
	Get(key string) (string, error)
	Keys(key string) ([]string, error)
	INCR(key string) error
	INCRBy(key string, num int64) error
	DECR(key string) error
	Expire(key string, ttl int) error
	SetEx(key string, value interface{}, ttl int) error
	SetNX(key string, value interface{}) (bool, error)
	Lpush(key string, value interface{}) error
	LPop(key string) (string, error)
	LRange(key string, start, end int64) ([]string, error)
	RPush(key string, value interface{}) error
	RPop(key string) (string, error)
	Llen(key string) (int64, error)
}
