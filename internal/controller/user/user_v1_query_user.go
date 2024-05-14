package user

import (
	"context"

	v1 "rme/api/user/v1"
	"rme/internal/service"
)

// QueryUser 查询用户基础信息
func (c *ControllerV1) QueryUser(ctx context.Context, req *v1.QueryUserReq) (res *v1.QueryUserRes, err error) {
	// 从登陆态中获取用户的user_id
	// 通过user_id查询db的到用户信息

	user, err := service.User().QueryUser(ctx, "user1") // todo: 如何从登陆态中获取信息?
	if err != nil {
		return nil, err
	}

	res = &v1.QueryUserRes{
		User: v1.UserBasic{
			NickName: user.Nickname,
			HeadUrl:  user.Headurl,
			Age:      int32(user.Age),
			Sex:      int32(user.Sex),
		},
	}

	return
}
