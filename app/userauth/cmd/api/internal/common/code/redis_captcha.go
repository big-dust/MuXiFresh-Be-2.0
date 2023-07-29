package code

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.Redis
	KeyPrefix   string
}

var redisClient *redis.Redis

// Set 实现 base64Captcha.Store interface 的 Set ⽅法
func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := 60 * conf.ExpireTime
	if err := s.RedisClient.Setex(s.KeyPrefix+key, value, ExpireTime); err != nil {
		return errors.New("⽆法存储图⽚验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get ⽅法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val, err := s.RedisClient.Get(key)
	if err != nil {
		return ""
	}
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify ⽅法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer && answer != ""
}
