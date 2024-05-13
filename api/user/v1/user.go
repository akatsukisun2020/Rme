package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 注意：所有关于用户的身份信息的字段，用户id/用户手机号/用户accessToken/用户refreshAccessToken都放在cookie中
// 如果端上不支持，则可以放到body中，由终端自己存储到本地storage

// UserBasic: 用户基础信息结构

type UserBasic struct {
	NickName string `json:"nick_name"`
	HeadUrl  string `json:"head_url"`
	Age      int32  `json:"age"`
	Sex      int32  `json:"sex"`
}

// QueryUser：用户的基本信息查询

type QueryUserReq struct {
	g.Meta `path:"/query_user" tags:"QueryUser" method:"get" summary:"查询用户信息"`
	// UserID string `json:"user_id"`  // 从登陆态中取
}

type QueryUserRes struct {
	User UserBasic `json:"user"` // 用户信息
}

// UpdateUser: 更新用户信息

type UpdateUserReq struct {
	g.Meta  `path:"/update_user" tags:"UpdateUser" method:"post" summary:"更新用户信息"`
	UFields []string  `json:"u_fields"` // 更新字段名
	User    UserBasic `json:"user"`     // 用户信息
}

type UpdateUserRes struct {
	ErrCode int32  `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误信息
}

// Login：用户登陆接口，在查不到用户信息的时候，直接进行注册(使用平台的默认头像和昵称).用户在注册成功之后，可以后面自行修改头像昵称等

type LoginReq struct {
	g.Meta    `path:"/login" tags:"Login" method:"post" summary:"用户登陆"`
	LoginType int32  `json:"login_type"` // 登陆方式, 0: wx登陆 1:手机号登陆
	Code      string `json:"code"`       // 登陆码
}

type LoginRes struct {
	ErrCode int32  `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误信息
}

// Authorize: 用户认证，主要适用于鉴权，看看这个用户是否登陆了。（有一些接口是需要有登陆态的）

type AuthorizeReq struct {
	g.Meta `path:"/authorize" tags:"Authorize" method:"post" summary:"用户认证"`
}

type AuthorizeRes struct {
	ErrCode int32  `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误信息
}

// RefreshToken: 刷新登陆态

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" tags:"RefreshToken" method:"post" summary:"刷新登陆台"`
}

type RefreshTokenRes struct {
	ErrCode int32  `json:"err_code"` // 错误码
	ErrMsg  string `json:"err_msg"`  // 错误信息
}
