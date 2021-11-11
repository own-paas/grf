package model

import (
	"github.com/sestack/grf/example/global"
)

type User struct {
	global.Model
	Name     string `json:"name" gorm:"uniqueIndex;not null;type:varchar(64);comment:用户名"` // 用户名
	Password string `json:"password" gorm:"not null;type:varchar(64);comment:密码"`          // 密码
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:电话号码"`                    // 电话号码
	Email    string `json:"email" gorm:"type:varchar(32);comment:邮箱"`                      // 邮箱
}
