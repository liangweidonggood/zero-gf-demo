# go-zero和go-frame结合使用

```
rpc启动
cd service/user/rpc
go run user.go -f etc/user.yaml

api启动
cd service/user/api
go run user.go -f etc/user-api.yaml
```

## 新建一个项目
```
go.mod文件

module zero-gf-demo

go 1.16

require (
	github.com/tal-tech/go-zero v1.1.7
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.29.1
	github.com/gogf/gf v1.16.2
)
go mod tidy
```

## 新建一个rpc
```
goctl template init
模板文件位于C:\Users\Administrator\.goctl

1.新建user.proto文件
2.生成rpc代码 
  cd user/rpc
  goctl rpc proto -src user.proto -dir .
3.生成dao代码
  cd zero-gf-demo
  gf gen dao -l "mysql:root:123@tcp(192.168.230.129:3306)/gozero-book"  -path ./service/user/rpc/internal  -t user
```
## 新建一个api
```
1.新建user.api
2.生成api代码
  cduser/api
  goctl api go -api user.api -dir .
```

## 生成一个model跟goframe的对比一下
```
cd user/rpc
goctl model mysql datasource -url="root:123@tcp(192.168.230.129:3306)/gozero-book" -table="user" -dir ./model -c
```

## 测试
```
curl -i -X POST http://localhost:8888/api/v1/user/one -H 'content-type: application/json' -d '{"id":1}'
```

## 新需求
```
1.用户登录 生成token
  前端第一次登录，调用授权api,生成accesstoken(2小时)和refreshtoken(7天)
  
2.token过期刷新token
  accesstoken过期refreshtoken没过期调用此接口返回新的accesstoken
  前端需要拦截token过期返回码来调用此接口
3.业务接口
  请求头需要带accesstoken
```
