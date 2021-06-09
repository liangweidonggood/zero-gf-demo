package svc

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"zero-gf-demo/service/user/rpc/internal/config"
	"zero-gf-demo/service/user/rpc/internal/dao"
	"zero-gf-demo/service/user/rpc/model"
)

type ServiceContext struct {
	Config  config.Config
	UserDao dao.UserDao

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Host: c.DBConfig.Host,
				Port: c.DBConfig.Port,
				User: c.DBConfig.User,
				Pass: c.DBConfig.Pass,
				Name: c.DBConfig.Name,
				Type: c.DBConfig.Type,
			},
		},
	})
	return &ServiceContext{
		Config: c,
		UserDao: dao.UserDao{
			M:     g.DB("default").Model("user").Safe(),
			DB:    g.DB("default"),
			Table: "user",
			Columns: dao.UserColumns{
				Id:         "id",
				Number:     "number",
				Name:       "name",
				Username:   "username",
				Password:   "password",
				Gender:     "gender",
				CreateTime: "create_time",
				UpdateTime: "update_time",
			},
		},
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
	}
}
