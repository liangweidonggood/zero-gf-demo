package logic

import (
	"context"

	"zero-gf-demo/service/user/rpc/internal/svc"
	"zero-gf-demo/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过用户名查询用户
func (l *GetUserByNameLogic) GetUserByName(in *user.UserNameReq) (*user.UserDetailRes, error) {
	u, err := l.svcCtx.UserModel.FindOneByUsername(in.Username)
	if err != nil {
		return nil, err
	}
	logx.Info(u)
	return &user.UserDetailRes{
		Id:       u.Id,
		Name:     u.Name,
		Number:   u.Number,
		Gender:   u.Gender,
		Username: u.Username,
		Password: u.Password,
	}, nil
}
