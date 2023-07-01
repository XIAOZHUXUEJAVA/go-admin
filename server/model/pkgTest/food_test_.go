// 自动生成模板Food
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Food 结构体
type Food struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:菜品名称;"`
}


// TableName Food 表名
func (Food) TableName() string {
  return "food"
}

