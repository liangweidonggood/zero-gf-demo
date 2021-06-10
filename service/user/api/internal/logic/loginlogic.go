package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
	"zero-gf-demo/common/errorx"
	"zero-gf-demo/service/user/rpc/model"
	"zero-gf-demo/service/user/rpc/userclient"

	"zero-gf-demo/service/user/api/internal/svc"
	"zero-gf-demo/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginRes, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserRPC.GetUserByName(l.ctx, &userclient.UserNameReq{
		Username: req.Username,
	})
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errorx.NewDefaultError("用户密码不正确")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginRes{
		Id:          userInfo.Id,
		Name:        userInfo.Name,
		Gender:      userInfo.Gender,
		Number:      userInfo.Number,
		Username:    userInfo.Username,
		AccessToken: jwtToken,
	}, nil
}

//根据安全码，过期时间，用户id生成一个token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
