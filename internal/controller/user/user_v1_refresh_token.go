package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "rme/api/user/v1"
	"rme/internal/consts"
	"rme/internal/service"
)

// curl -H "Content-Type: application/json" -H "RmeUserid: 363bea77-bc7d-4be0-81f5-27f0aafb748a" -H "RmeRefreshToken: JOgm2GQvcZfeWNwmA9pNP2DRRDK+hYqdFle8aqJRseEB+J6BZUKFqETZcO0UPtpB+RskBMolm+5J/RX57yMULRvihjylvokMajFZPT2DFbyQirO594TnyQs04WOSefYB8MnuGfGd67PhUBxKW9Yin87clHYzdc3dgigrjzo2dtDhxZ9oy3XJlWbSQg==" -X POST  -d '{}' http://127.0.0.1:8000/user/refresh_token

// RefreshToken 登陆态过期的时候使用，用来刷新登陆态
func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = new(v1.RefreshTokenRes)
	userID := g.RequestFromCtx(ctx).GetHeader(consts.HttpHeader_UserId)
	refreshToken := g.RequestFromCtx(ctx).GetHeader(consts.HttpHeader_RefreshToken)

	loginInfo, err := service.User().RefreshToken(ctx, userID, refreshToken)
	if err != nil {
		g.Log().Errorf(ctx, "RefreshToken failed, userid:%s", userID)
		res.ErrCode = consts.ErrCode_RefreshToken
		res.ErrMsg = "登陆态刷新失败"
		return
	}

	// 无登陆态或者已过期，提示重新登录
	if loginInfo == nil {
		g.Log().Infof(ctx, "RefreshToken failed, need login, userid:%s", userID)
		res.ErrCode = consts.ErrCode_NeedLogin
		res.ErrMsg = "请重新登录"
		return res, nil
	}
	// 成功刷新了登陆态
	res.LoginInfo = v1.LoginInfo{
		UserID:                loginInfo.UserId,
		AccessToken:           loginInfo.AccessToken,
		RereshToken:           loginInfo.RefreshToken,
		AccessTokenExpireTime: loginInfo.AccessTokenExpireTime,
	}

	return res, nil
}
