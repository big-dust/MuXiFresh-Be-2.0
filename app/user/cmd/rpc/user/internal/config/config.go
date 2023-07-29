package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MongoConf struct {
		URL string
		DB  string
	}
	Limit int64
}
