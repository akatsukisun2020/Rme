package model

// 此协议负责定义service层的协议

// User is the golang structure for table user.
type User struct {
	Userid       string `json:"userid"`         // 用户id
	Age          uint   `json:"age"`            // 用户年纪
	Sex          uint   `json:"sex"`            // 用户性别
	Headurl      string `json:"headurl"`        // 用户头像
	Nickname     string `json:"nickname"`       // 用户昵称
	Mobile       string `json:"mobile"`         // 电话号码
	Email        string `json:"email"`          // 邮件地址
	UserSecretId string `json:"user_secret_id"` // 用户密钥id(单用户唯一不变id)
}

// LoginInfo 用户登陆信息, redis存储 kv结构，key:"rme_logininfo_${userid}"
type LoginInfo struct {
	UserId                 string `json:"user_id"`                   // 用户id
	UserSecretId           string `json:"user_secret_id"`            // 用户密钥id
	LoginType              int    `json:"login_type"`                // 登陆类型 0:ios 1:安卓 2:web
	AccessToken            string `json:"access_token"`              // token
	RefreshToken           string `json:"refresh_token"`             // 刷新token
	AccessTokenExpireTime  int64  `json:"access_token_expire_time"`  // access token生效时间（用于判定过期)
	RefreshTokenExpireTime int64  `json:"refresh_token_expire_time"` // refresh token生效时间（用于判定过期)
	Mac                    string `json:"mac"`                       // mac地址
}
