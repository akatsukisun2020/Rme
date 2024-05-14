package user

import (
	"context"

	v1 "rme/api/user/v1"
	"rme/internal/model"
	"rme/internal/service"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	// todo: 直接设置一个数据进去玩玩

	err = service.User().CreateUser(ctx, &model.User{
		Age:      18,
		Sex:      0,
		Nickname: "bigsun",
		Mobile:   "123456",
	})

	return &v1.UpdateUserRes{}, nil
}
