package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "rme/api/user/v1"
	"rme/internal/consts"
	"rme/internal/service"
)

// curl -H "Content-Type: application/json" -H "RmeUserid: 41fc5a19-df59-4361-8865-65e6d55c721c" -H "RmeAccessToken: 8Pn2Pm1k1Ah5eqQJQpcKtw+1xvz4DdNO/2Fp71xzbBzNn6gFyWAdDZzLLiIYVVk+d1ioFRqj+a9s6X6WNpvPiv6pTV+Lxis3Smzzw7c+mPj1zDaDi7QdIbFs+ywkcwiNopmYlH4WuyMASv251QhvISCmj4c2DopHjDfLzwpO9fvrM0lXmE9S82Ieag==" -X POST  -d "{}" http://127.0.0.1:8000/user/verify

// Verify 登录态校验(如何在每个接口都调用?拦截器中实现？)
func (c *ControllerV1) Verify(ctx context.Context, req *v1.VerifyReq) (res *v1.VerifyRes, err error) {
	res = new(v1.VerifyRes)
	userID := g.RequestFromCtx(ctx).GetHeader(consts.HttpHeader_UserId)
	accessToken := g.RequestFromCtx(ctx).GetHeader(consts.HttpHeader_AccessToken)

	g.Log().Debug(ctx, "aaaaaaaaaaaa1111aaaaaa33")
	valid, err := service.User().Verify(ctx, userID, accessToken)
	if err != nil {
		g.Log().Errorf(ctx, "Verify failed, userid:%s", userID)
		res.ErrCode = consts.ErrCode_Verify
		res.ErrMsg = "登陆态校验失败"
		return
	}

	g.Log().Debug(ctx, "aaaaaaaaaaaa1111aaaaaa")
	if !valid {
		res.ErrCode = consts.ErrCode_NotLogin
		res.ErrMsg = "未登陆或登陆态过期"
		return
	}

	res.ErrMsg = "登陆态校验通过"
	g.Log().Debug(ctx, "aaaaaaaaaaaaaaaaaa")

	return res, nil
}
