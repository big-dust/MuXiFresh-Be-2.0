package config

import (
	"MuXiFresh-Be-2.0/common/email"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	FormConf  zrpc.RpcClientConf
	EmailConf email.SenderConf
	JwtAuth   struct {
		AccessSecret string
		AccessExpire int64
	}
	MongoDBConf struct {
		URL string
		DB  string
	}
}
