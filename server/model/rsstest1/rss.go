// 自动生成模板Rss
package rsstest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// Rss 结构体
type Rss struct {
      global.GVA_MODEL
      Title  string `json:"rss_title" form:"rss_title" gorm:"column:title;comment:标题;size:255;"`
      Description  string `json:"rss_description" form:"rss_description" gorm:"column:description;comment:描述;size:255;"`
}


// TableName Rss 表名
func (Rss) TableName() string {
  return "rss"
}

