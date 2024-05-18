// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"rme/api/user/v1"
)

type IUserV1 interface {
	QueryUser(ctx context.Context, req *v1.QueryUserReq) (res *v1.QueryUserRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Verify(ctx context.Context, req *v1.VerifyReq) (res *v1.VerifyRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
}
