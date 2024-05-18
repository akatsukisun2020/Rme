package user

import (
	"context"
	"rme/internal/consts"
	"rme/internal/dao"
	"rme/internal/model"
	"rme/internal/model/do"
	"rme/internal/model/entity"
	"rme/internal/service"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

// New 创建User对象
func New() service.IUser {
	return &sUser{}
}

// GenUserID 生成用户id
func GenUserID() string {
	return uuid.New().String()
}

// InsertOrUpdateUser 创建用户, 存在则更新，不存在则插入
func (s *sUser) InsertOrUpdateUser(ctx context.Context, user *model.User) (err error) {
	if user.Userid == "" {
		user.Userid = GenUserID() // 为新用户生成用户id
	}

	res, err := dao.User.Ctx(ctx).Save(user)
	if err != nil {
		g.Log().Errorf(ctx, "dao.Save error, user:%v, err:%v", user, err)
		return err
	}

	g.Log().Infof(ctx, "InsertOrUpdateUser success, user:%v, res:%v", user, res)
	return nil
}

// RemoveUser 删除用户
func (s *sUser) RemoveUser(ctx context.Context, userID string) (err error) {
	if userID == "" {
		g.Log().Error(ctx, "userid is emtpy")
		return gerror.New("userid is empty")
	}

	_, err = dao.User.Ctx(ctx).Delete(do.User{
		Userid: userID,
	})
	if err != nil {
		g.Log().Errorf(ctx, "dao.Delete error, userid:%s, err:%v", userID, err)
		return err
	}

	return nil
}

// QueryUser 查询用户
func (s *sUser) QueryUser(ctx context.Context, userID string) (res *entity.User, err error) {
	if userID == "" {
		g.Log().Error(ctx, "userid is emtpy")
		return nil, gerror.New("userid is empty")
	}

	// 根据userid查询用户信息
	if err = dao.User.Ctx(ctx).Where(do.User{
		Userid: userID,
	}).Scan(&res); err != nil {
		g.Log().Errorf(ctx, "dao.Scan error, userid:%s, err:%s", userID, err.Error())
	}

	return
}

// QueryUserByLw 根据登陆方式查询用户
func (s *sUser) QueryUserByLw(ctx context.Context, loginType int32, contact string) (res *entity.User, err error) {
	if contact == "" {
		g.Log().Error(ctx, "contact is emtpy")
		return nil, gerror.New("contact is empty")
	}

	doUser := do.User{}
	switch loginType {
	case consts.LoginType_Mobile:
		doUser.Mobile = contact
	default:
		g.Log().Errorf(ctx, "no such loginType, loginType:%d, contact:%s", loginType, contact)
		return nil, gerror.New("no such loginType")
	}

	// 根据userid查询用户信息
	if err = dao.User.Ctx(ctx).Where(doUser).Scan(&res); err != nil {
		g.Log().Errorf(ctx, "dao.Scan error, loginType:%d, contact:%s, err:%v", loginType, contact, err)
		return nil, err
	}

	return
}

// UpdateUser 更新用户 能否指定字段进行更新（TODO）
func (s *sUser) UpdateUser(ctx context.Context, fields []string, user *model.User) (err error) {

	if user.Userid == "" {
		g.Log().Errorf(ctx, "UpdateUser param error, userid is empty, user:%v", user)
		return gerror.New("UpdateUser param error")
	}

	doUser := do.User{
		Userid: user.Userid,
	}
	for _, v := range fields {
		switch v {
		case "nick_name":
			doUser.Nickname = user.Nickname
		case "head_url":
			doUser.Headurl = user.Headurl
		case "age":
			doUser.Age = user.Age
		case "sex":
			doUser.Sex = user.Sex
		default:
			g.Log().Errorf(ctx, "fields error, userid:%s, fileds:%v", user.Userid, fields)
			return gerror.New("UpdateUser fields error")
		}
	}

	_, err = dao.User.Ctx(ctx).Data(doUser).Where("userid", user.Userid).Update()
	if err != nil {
		g.Log().Errorf(ctx, "UpdateUser error, userid:%s, fileds:%v", user.Userid, fields)
		return err
	}

	return nil
}

// Verify 登陆态校验
func (s *sUser) Verify(ctx context.Context, userID, accessToken string) (valid bool, err error) {
	if userID == "" || accessToken == "" {
		return false, nil
	}

	// redis中读取必要信息
	info, err := dao.NewLoginInfoRedisClient().Get(ctx, userID)
	if err != nil {
		g.Log().Errorf(ctx, "redis get error, err:%v", err)
		return false, err
	}

	// 校验必要的身份id和token
	if info.UserId != userID || info.AccessToken != accessToken {
		g.Log().Debugf(ctx, "user not login, userid:%s, accessToken:%s, redisUserId:%s, redisAccessToken:%s",
			userID, accessToken, info.UserId, info.AccessToken)
		return false, nil
	}

	// 解析token，校验accessToken时效性和正确性
	tokener := &Tokener{
		Token:        accessToken,
		UserSecretID: info.UserSecretId,
	}
	tokener.ParseToken(ctx)
	nowts := time.Now().Unix()
	if tokener.UserID != userID || tokener.UserSecretID != info.UserSecretId ||
		tokener.TokenType != consts.TokenType_AccessToken ||
		tokener.ExpireTime <= nowts {
		return false, nil
	}

	// 至此，表示登录态校验通过
	return true, nil
}

// Login 登陆授权
func (s *sUser) Login(ctx context.Context, loginType int, code string) (loginInfo *model.LoginInfo, err error) {
	// 校验登陆是否还有效，如果有效，则不要再次进行登陆了？【但是，这个逻辑需要引入端】，双端登陆？是否能够支持双端登陆，or一端登陆的时候，挤掉另一端？
	// code去重拦截&ip拦截等，避免浪费流量（money） // TODO: 后续有时间在做这个.
	// oauth2.0协议，从第三方获取手机号or其他身份字段信息 // TODO: 需要看采用何种方式 ，手机还是wx
	// 查询用户信息表，是否存在此用户，不存在则创建，存在，则返回认证登陆态
	// 返回uid&token等信息

	if loginType != consts.LoginType_WX && loginType != consts.LoginType_Mobile || code == "" {
		g.Log().Errorf(ctx, "login param error, loginType:%d, code:%s", loginType, code)
		return nil, gerror.New("login param error")
	}

	// mock 代码
	contact := "123456111331"
	user, err := s.QueryUserByLw(ctx, int32(loginType), contact)
	if err != nil {
		g.Log().Errorf(ctx, "QueryUserByLw error, err:%v", err)
		return nil, err
	}

	// 生成新的用户信息
	loginInfo = new(model.LoginInfo)
	if user == nil { // todo： 这一行和上面检测，理论上需要加锁，否则可能搞出2个useid
		muser := new(model.User)
		muser.UserSecretId = uuid.NewString()
		if loginType == consts.LoginType_Mobile {
			muser.Mobile = contact
		} else {
			// todo
		}
		if err = s.InsertOrUpdateUser(ctx, muser); err != nil {
			g.Log().Errorf(ctx, "InsertOrUpdateUser error, err:%v", err)
			return nil, err
		}

		loginInfo.UserId = muser.Userid
		loginInfo.UserSecretId = muser.UserSecretId
		loginInfo.LoginType = loginType
	} else {
		loginInfo.UserId = user.Userid
		loginInfo.UserSecretId = user.Usersecretid
		loginInfo.LoginType = loginType
	}

	nowMs := time.Now().Unix()
	tokenValidityPeriod, _ := g.Cfg().Get(ctx, "custom.token.token_validity_period")
	refreshValidityPeriod, _ := g.Cfg().Get(ctx, "custom.token.refresh_validity_period")
	loginInfo.AccessTokenExpireTime = nowMs + tokenValidityPeriod.Int64()
	loginInfo.RefreshTokenExpireTime = nowMs + refreshValidityPeriod.Int64()

	// 生成accessToken和refreshToken
	accessTokener := &Tokener{
		UserID:       loginInfo.UserId,
		UserSecretID: loginInfo.UserSecretId,
		TokenType:    consts.TokenType_AccessToken,
		ExpireTime:   loginInfo.AccessTokenExpireTime,
	}
	if err = accessTokener.GenerateToken(ctx); err != nil {
		g.Log().Errorf(ctx, "GenerateToken error, accessTokener:%v, err:%v", accessTokener, err)
	}

	refreshTokener := &Tokener{
		UserID:       loginInfo.UserId,
		UserSecretID: loginInfo.UserSecretId,
		TokenType:    consts.TokenType_AccessToken,
		ExpireTime:   loginInfo.AccessTokenExpireTime,
	}
	if err = refreshTokener.GenerateToken(ctx); err != nil {
		g.Log().Errorf(ctx, "GenerateToken error, accessTokener:%v, err:%v", accessTokener, err)
	}
	loginInfo.AccessToken = accessTokener.Token
	loginInfo.RefreshToken = refreshTokener.Token

	// 覆盖redis登陆态
	if err = dao.NewLoginInfoRedisClient().Set(ctx, loginInfo); err != nil {
		g.Log().Errorf(ctx, "redis set error, err:%v", err)
		return nil, err
	}

	return loginInfo, nil
}

// RefreshToken 刷新登录态
func (s *sUser) RefreshToken(ctx context.Context, userID, refreshToken string) (loginInfo *model.LoginInfo, err error) {
	// 频率等限制，ip限制等 TODO
	// 校验refreshToken，看看是否过期，userid是否对应得上等
	// 校验通过之后，生成新的token，原先的token作废

	if userID == "" || refreshToken == "" {
		g.Log().Errorf(ctx, "RefreshToken param error, userid:%s, refreshToken:%s", userID, refreshToken)
		return nil, gerror.New("RefreshToken param error")
	}

	// 从redis中，先获取登陆态信息
	redisLoginInfo, err := dao.NewLoginInfoRedisClient().Get(ctx, userID)
	if err != nil {
		g.Log().Errorf(ctx, "redis set error, err:%v", err)
		return nil, err
	}
	if redisLoginInfo == nil { // 表明根本没有登录，或者早就过期了。提示需要重新登录
		g.Log().Infof(ctx, "user need login, userid:%s", userID)
		return nil, nil // 返回nil，需要重新登录
	}

	// 校验登陆态
	refreshTokener := &Tokener{
		UserSecretID: redisLoginInfo.UserSecretId,
		Token:        redisLoginInfo.RefreshToken,
	}
	if err = refreshTokener.ParseToken(ctx); err != nil {
		g.Log().Errorf(ctx, "GenerateToken error, refreshTokener:%v, err:%v", refreshTokener, err)
		return nil, err
	}
	if refreshTokener.UserID != userID || refreshTokener.Token != refreshToken ||
		refreshTokener.ExpireTime <= time.Now().Unix() {
		g.Log().Errorf(ctx, "GenerateToken error, refreshToken is not usable, userid:%s, refreshToken:%s",
			userID, refreshToken)
		return nil, nil // 返回nil，需要重新登录
	}

	// 重新生成新的token
	nowMs := time.Now().Unix()
	tokenValidityPeriod, _ := g.Cfg().Get(ctx, "custom.token.token_validity_period")
	refreshValidityPeriod, _ := g.Cfg().Get(ctx, "custom.token.refresh_validity_period")
	loginInfo = &model.LoginInfo{
		UserId:       userID,
		UserSecretId: redisLoginInfo.UserSecretId,
		LoginType:    redisLoginInfo.LoginType,
	}
	loginInfo.AccessTokenExpireTime = nowMs + tokenValidityPeriod.Int64()
	loginInfo.RefreshTokenExpireTime = nowMs + refreshValidityPeriod.Int64()

	// 生成accessToken和refreshToken
	accessTokener := &Tokener{
		UserID:       loginInfo.UserId,
		UserSecretID: loginInfo.UserSecretId,
		TokenType:    consts.TokenType_AccessToken,
		ExpireTime:   loginInfo.AccessTokenExpireTime,
	}
	if err = accessTokener.GenerateToken(ctx); err != nil {
		g.Log().Errorf(ctx, "GenerateToken error, accessTokener:%v, err:%v", accessTokener, err)
	}
	refreshTokener = &Tokener{
		UserID:       loginInfo.UserId,
		UserSecretID: loginInfo.UserSecretId,
		TokenType:    consts.TokenType_AccessToken,
		ExpireTime:   loginInfo.AccessTokenExpireTime,
	}
	if err = refreshTokener.GenerateToken(ctx); err != nil {
		g.Log().Errorf(ctx, "GenerateToken error, accessTokener:%v, err:%v", accessTokener, err)
	}
	loginInfo.AccessToken = accessTokener.Token
	loginInfo.RefreshToken = refreshTokener.Token

	// 覆盖redis登陆态
	if err = dao.NewLoginInfoRedisClient().Set(ctx, loginInfo); err != nil {
		g.Log().Errorf(ctx, "redis set error, err:%v", err)
		return nil, err
	}

	return loginInfo, nil
}
