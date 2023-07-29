package ctxData

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	CtxKeyJwtEmail  = "jwtEmail"
	CtxKeyJwtUserID = "jwtUserId"
)

func GetEmailFromCtx(ctx context.Context) string {
	email, ok := ctx.Value(CtxKeyJwtEmail).(string)
	if !ok {
		logx.WithContext(ctx).Errorf("GetEmailFromCtx failed")
	}
	return email
}

func GetUserIdFromCtx(ctx context.Context) string {
	userID, ok := ctx.Value(CtxKeyJwtUserID).(string)
	if !ok {
		logx.WithContext(ctx).Errorf("GetEmailFromCtx failed")
	}
	return userID
}
