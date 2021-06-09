package logic

import (
	"context"
	"zero-gf-demo/service/user/rpc/userclient"

	"zero-gf-demo/service/user/api/internal/svc"
	"zero-gf-demo/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) OneLogic {
	return OneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OneLogic) One(req types.UserOneReq) (*userclient.UserInfoReply, error) {
	return l.svcCtx.UserRPC.GetUser(l.ctx, &userclient.IdReq{
		Id: req.Id,
	})
}
