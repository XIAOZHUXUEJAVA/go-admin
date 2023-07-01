// 自动生成模板OSUser
package usertest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// OSUser 结构体
type OSUser struct {
      global.GVA_MODEL
      Username  string `json:"user_name" form:"user_name" gorm:"column:username;comment:用户名;size:255;"`
      Password  string `json:"user_password" form:"user_password" gorm:"column:password;comment:密码;size:255;"`
}


// TableName OSUser 表名
func (OSUser) TableName() string {
  return "user"
}

