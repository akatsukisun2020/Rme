package user

import (
	"context"
	"rme/internal/dao"
	"rme/internal/model"
	"rme/internal/model/do"
	"rme/internal/model/entity"
	"rme/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// CreateUser 创建用户
func (s *sUser) CreateUser(ctx context.Context, user *model.User) (err error) {
	if user.Userid != "" {
		return gerror.Newf("userid is not empty, userid:%s", user.Userid)
	}

	user.Userid = genUserID() // 为新用户生成用户id
	_, err = dao.User.Ctx(ctx).Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *sUser) RemoveUser(ctx context.Context, userID string) (err error) {

	return nil
}

func (s *sUser) QueryUser(ctx context.Context, userID string) (res *entity.User, err error) {
	if userID == "" {
		return nil, gerror.New("userid is empty")
	}

	// 根据userid查询用户信息
	err = dao.User.Ctx(ctx).Where(do.User{
		Userid: userID,
	}).Scan(&res)

	return
}

func (s *sUser) UpdateUser(ctx context.Context, user *model.User) (err error) {

	return nil
}

func genUserID() string {
	return uuid.New().String()
}
