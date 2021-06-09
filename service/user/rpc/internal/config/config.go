package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	//加入goframe中mysql配置
	DBConfig struct {
		Host string
		Port string
		User string
		Pass string
		Name string
		Type string
	}

	Mysql struct {
		DataSource string
	}
	CacheRedis cache.ClusterConf
}
