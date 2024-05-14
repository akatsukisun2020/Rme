package model

// 此协议负责定义service层的协议

// User is the golang structure for table user.
type User struct {
	Userid   string `json:"userid"`   // 用户id
	Age      uint   `json:"age"`      // 用户年纪
	Sex      uint   `json:"sex"`      // 用户性别
	Headurl  string `json:"headurl"`  // 用户头像
	Nickname string `json:"nickname"` // 用户昵称
	Mobile   string `json:"mobile"`   // 电话号码
	Email    string `json:"email"`    // 邮件地址
}
