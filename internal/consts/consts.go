package consts

const ( // TokenType 定义
	TokenType_AccessToken  = 0 // accetoken 类型
	TokenType_RefreshToken = 1 // refresh_token 类型
)

const (
	LoginType_WX     = 1 // 微信登陆
	LoginType_Mobile = 2 // 手机登陆
)

const ( // http头部字段定义
	HttpHeader_UserId          = "RmeUserid"          // 用户id
	HttpHeader_AccessToken     = "RmeAccessToken"     // token码
	HttpHeader_RefreshToken    = "RmeRefreshToken"    // 刷新码
	HttpHeader_RefreshInterval = "RmeRefreshInterval" // 刷新间隔
)

const ( // 返回码定义
	ErrCode_Success = 0

	ErrCode_Verify       = -100 // 登陆态校验失败
	ErrCode_Login        = -101 // 登陆失败
	ErrCode_RefreshToken = -102 // 登陆态校验失败

	ErrCode_NotLogin  = 100 // 未登陆
	ErrCode_NeedLogin = 101 // 需要登录
)
