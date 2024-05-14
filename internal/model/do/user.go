// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	Id       interface{} // 自增id
	Userid   interface{} // 用户id
	Age      interface{} // 用户年纪
	Sex      interface{} // 用户性别
	Headurl  interface{} // 用户头像
	Nickname interface{} // 用户昵称
	Mobile   interface{} // 电话号码
	Email    interface{} // 邮件地址
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
	DeleteAt *gtime.Time // Deleted Time
}
