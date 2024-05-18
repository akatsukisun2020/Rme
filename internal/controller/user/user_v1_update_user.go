package user

import (
	"context"

	v1 "rme/api/user/v1"
	"rme/internal/consts"
	"rme/internal/model"
	"rme/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// curl -H "Content-Type: application/json" -H "RmeUserid: 41fc5a19-df59-4361-8865-65e6d55c721c"  -X POST  -d '{"u_fields":["nick_name", "head_url"], "user":{"nick_name":"A++", "head_url":"www.baidu.com"}}' http://127.0.0.1:8000/user/update_user

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	res = new(v1.UpdateUserRes)
	userID := g.RequestFromCtx(ctx).GetHeader(consts.HttpHeader_UserId)

	if len(req.UFields) == 0 || userID == "" {
		g.Log().Errorf(ctx, "UpdateUser param error, userid:%s, req:%v", userID, req)
		res.ErrCode = consts.ErrCode_UpdateUser
		res.ErrMsg = "更新用户信息参数错误"
		return
	}

	mUser := &model.User{
		Userid:   userID,
		Age:      uint(req.User.Age),
		Sex:      uint(req.User.Sex),
		Headurl:  req.User.HeadUrl,
		Nickname: req.User.NickName,
	}

	if err = service.User().UpdateUser(ctx, req.UFields, mUser); err != nil {
		g.Log().Errorf(ctx, "UpdateUser error, euserid:%s, err:%v", userID, err)
		res.ErrCode = consts.ErrCode_UpdateUser
		res.ErrMsg = "更新用户信息失败"
		return
	}

	return
}
