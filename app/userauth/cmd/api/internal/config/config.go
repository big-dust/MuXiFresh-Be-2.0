package config

import (
	"MuXiFresh-Be-2.0/common/code"
	"MuXiFresh-Be-2.0/common/email"
	"MuXiFresh-Be-2.0/common/tube"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AccountCenterConf zrpc.RpcClientConf
	EmailConf         email.SenderConf
	Oss               tube.Qiniu
	CaptchaConf       code.Conf
	RedisConf         redis.RedisConf
	JwtAuthChPass     struct {
		AccessSecret string
		AccessExpire int64
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
