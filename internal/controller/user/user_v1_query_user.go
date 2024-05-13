package user

import (
	"context"

	v1 "rme/api/user/v1"
)

// QueryUser 查询用户基础信息
func (c *ControllerV1) QueryUser(ctx context.Context, req *v1.QueryUserReq) (res *v1.QueryUserRes, err error) {
	// 从登陆态中获取用户的user_id
	// 通过user_id查询db的到用户信息

	res = &v1.QueryUserRes{
		User: v1.UserBasic{
			NickName: "goodman",
			HeadUrl:  "www.baidu.com",
			Age:      18,
			Sex:      0,
		},
	}

	return
}
