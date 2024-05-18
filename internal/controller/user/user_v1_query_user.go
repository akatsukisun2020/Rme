package user

import (
	"context"
	"fmt"

	v1 "rme/api/user/v1"
	"rme/internal/consts"
	"rme/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// 测试命令：
// curl -H "Content-Type: application/json" -H "RmeUserid: 0384c046-b157-4c74-9f77-d0fb7f7e5675" http://127.0.0.1:8000/user/query_user

// QueryUser 查询用户基础信息
func (c *ControllerV1) QueryUser(ctx context.Context, req *v1.QueryUserReq) (res *v1.QueryUserRes, err error) {
	// 从登陆态中获取用户的user_id
	userID := g.RequestFromCtx(ctx).GetHeader(consts.HttpHeader_UserId)
	user, err := service.User().QueryUser(ctx, userID)
	if err != nil {
		g.Log().Errorf(ctx, "service.QueryUser error, userid:%s, err:%v", userID, err)
		return nil, err
	}

	// fixme: 可能查询db为空
	if user == nil {
		g.Log().Errorf(ctx, "user is not existed, userid:%s", userID)
		return nil, fmt.Errorf("user is not existed")
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
