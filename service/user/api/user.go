package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"zero-gf-demo/common/errorx"
	"zero-gf-demo/service/user/api/internal/config"
	"zero-gf-demo/service/user/api/internal/handler"
	"zero-gf-demo/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.OkJson(w, errorx.CodeError{Code: http.StatusUnauthorized, Msg: err.Error()})
	}))
	defer server.Stop()

	//跨域
	//server.Use(middleware.NewCorsMiddleware().Handle)
	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		fmt.Print(err)
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, errorx.CodeError{Code: http.StatusInternalServerError, Msg: e.Error()}
		}
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
