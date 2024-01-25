package code

import (
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// EmailCodeExpired 邮箱验证码过期时间
var EmailCodeExpired int

// Conf captcha 配置
type Conf struct {
	Height           int
	Width            int
	Length           int
	Maxskew          float64
	Dotcount         int
	ExpireTime       int
	DebugExpireTime  int
	TestingKey       string
	EmailCodeExpired int
}

var conf Conf

func Load(c *Conf, redisConf redis.RedisConf) {
	copier.Copy(&conf, c)
	redisClient = redis.MustNewRedis(redisConf)
	EmailCodeExpired = c.EmailCodeExpired
}
