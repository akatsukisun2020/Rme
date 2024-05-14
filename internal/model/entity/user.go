// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       uint        `json:"id"       orm:"id"        ` // 自增id
	Userid   string      `json:"userid"   orm:"userid"    ` // 用户id
	Age      uint        `json:"age"      orm:"age"       ` // 用户年纪
	Sex      uint        `json:"sex"      orm:"sex"       ` // 用户性别
	Headurl  string      `json:"headurl"  orm:"headurl"   ` // 用户头像
	Nickname string      `json:"nickname" orm:"nickname"  ` // 用户昵称
	Mobile   string      `json:"mobile"   orm:"mobile"    ` // 电话号码
	Email    string      `json:"email"    orm:"email"     ` // 邮件地址
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" ` // Created Time
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" ` // Updated Time
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" ` // Deleted Time
}
