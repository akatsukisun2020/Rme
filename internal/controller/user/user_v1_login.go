package user

import (
	"context"

	v1 "rme/api/user/v1"
	"rme/internal/consts"
	"rme/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// curl -H "Content-Type: application/json" -H "RmeUserid: 0384c046-b157-4c74-9f77-d0fb7f7e5675" -H "RmeAccessToken: 0384c046-b157-4c74-9f77-d0fb7f7e5675" -X POST  -d "{"login_type":2, "code":"AAA"}" http://127.0.0.1:8000/user/login

// Login 用户登陆，返回核心登陆态
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	loginInfo, err := service.User().Login(ctx, int(req.LoginType), req.Code)
	if err != nil {
		g.Log().Errorf(ctx, "service.Login error, req:%v, err:%v", req, err)
		res = &v1.LoginRes{
			ErrCode: consts.ErrCode_Login,
			ErrMsg:  "登陆失败",
		}
		return res, err
	}

	res = &v1.LoginRes{
		LoginInfo: v1.LoginInfo{
			UserID:                loginInfo.UserId,
			AccessToken:           loginInfo.AccessToken,
			RereshToken:           loginInfo.RefreshToken,
			AccessTokenExpireTime: loginInfo.AccessTokenExpireTime,
		},
	}
	return res, nil
}
