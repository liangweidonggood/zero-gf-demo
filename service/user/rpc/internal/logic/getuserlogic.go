package logic

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
	"zero-gf-demo/service/user/rpc/internal/svc"
	"zero-gf-demo/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	logx.Info("GetUser func in.....")
	var rep user.UserInfoReply
	db := g.DB()
	db.SetDebug(true)
	err := l.svcCtx.UserDao.Cache(time.Hour, "user-"+gconv.String(in.Id)).Scan(&rep, "id", in.Id)
	if err != nil {
		return nil, err
	}
	return &rep, nil
	//go-zero内部实现
	/*	one, err := l.svcCtx.UserModel.FindOne(in.Id)
		if err != nil {
			return nil, err
		}
		return &user.UserInfoReply{
			Id:     one.Id,
			Name:   one.Name,
			Number: one.Number,
			Gender: one.Gender,
		}, err*/

}
