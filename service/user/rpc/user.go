package main

import (
	"flag"
	"fmt"
	"zero-gf-demo/service/user/rpc/internal/config"
	"zero-gf-demo/service/user/rpc/internal/server"
	"zero-gf-demo/service/user/rpc/internal/svc"
	"zero-gf-demo/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//db := g.DB()
	//db.SetDebug(true)
	//r, _ := db.Table("user").Cache(time.Hour, "vip-user").Where("id", 1).One()
	//g.Log().Print(r.Map())
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
