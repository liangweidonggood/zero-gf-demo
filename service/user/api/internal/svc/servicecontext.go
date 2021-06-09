package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"zero-gf-demo/service/user/api/internal/config"
	"zero-gf-demo/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	//配置rpc
	UserRPC userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
