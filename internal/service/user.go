package service

import (
	"context"

	"rme/internal/model"
	"rme/internal/model/entity"
)

// IUser user接口
type IUser interface {
	// 资料管理相关接口
	InsertOrUpdateUser(ctx context.Context, user *model.User) (err error)       // 增
	RemoveUser(ctx context.Context, userID string) (err error)                  // 删
	UpdateUser(ctx context.Context, user *model.User) (err error)               // 修
	QueryUser(ctx context.Context, userID string) (res *entity.User, err error) // 查

	// 登陆态管理相关接口
	// redis操作/cookie操作/外部http协议访问
	// todo: 思考，是否应该进一步区分新的类？

	Verify(ctx context.Context, userID, accessToken string) (valid bool, err error)                        // 登录台校验
	Login(ctx context.Context, loginType int, code string) (loginInfo *model.LoginInfo, err error)         // 登陆授权
	RefreshToken(ctx context.Context, userID, refreshToken string) (loginInfo *model.LoginInfo, err error) // 刷新登陆态
}

var localUser IUser

// User 构造器
func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

// RegisterUser 注册user
func RegisterUser(i IUser) {
	localUser = i
}

// [FIXME] 关于web服务协议分层的理解:
// api协议层
// control层: 则是根据若干不同的service来完成自身的大的业务逻辑
// service协议层
// service层：也是设计到业务逻辑，需要聚焦于子领域; 不能仅仅封装db
// logic层：负责service层的"多态“实现
// dao协议层
// dao层
