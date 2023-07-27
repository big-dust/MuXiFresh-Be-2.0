package logic

import (
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/common/code"
	"MuXiFresh-Be-2.0/common/ctxData"
	"MuXiFresh-Be-2.0/common/globalKey"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/svc"
	"MuXiFresh-Be-2.0/app/userauth/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthPassChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthPassChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthPassChangeLogic {
	return &AuthPassChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthPassChangeLogic) AuthPassChange(req *types.AuthPassChangeReq) (resp *types.AuthPassChangeResp, err error) {
	if ok := code.VerifyEm(globalKey.AuthChPass, req.Email, req.VerifyCode); !ok {
		return nil, fmt.Errorf("verify code failed")
	}
	//gen auth token
	authChPassToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuthChPass.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuthChPass.AccessExpire, req.Email)
	if err != nil {
		return nil, err
	}
	return &types.AuthPassChangeResp{
		AuthChPassToken: authChPassToken,
	}, nil
}

func (l *AuthPassChangeLogic) getJwtToken(secretKey string, iat, seconds int64, email string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxData.CtxKeyJwtEmail] = email
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
