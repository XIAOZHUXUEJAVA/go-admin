package rsstest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rsstest1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    rsstest1Req "github.com/flipped-aurora/gin-vue-admin/server/model/rsstest1/request"
)

type RssService struct {
}

// CreateRss 创建Rss记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssService *RssService) CreateRss(rss *rsstest1.Rss) (err error) {
	err = global.GVA_DB.Create(rss).Error
	return err
}

// DeleteRss 删除Rss记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssService *RssService)DeleteRss(rss rsstest1.Rss) (err error) {
	err = global.GVA_DB.Delete(&rss).Error
	return err
}

// DeleteRssByIds 批量删除Rss记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssService *RssService)DeleteRssByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]rsstest1.Rss{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRss 更新Rss记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssService *RssService)UpdateRss(rss rsstest1.Rss) (err error) {
	err = global.GVA_DB.Save(&rss).Error
	return err
}

// GetRss 根据id获取Rss记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssService *RssService)GetRss(id uint) (rss rsstest1.Rss, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rss).Error
	return
}

// GetRssInfoList 分页获取Rss记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssService *RssService)GetRssInfoList(info rsstest1Req.RssSearch) (list []rsstest1.Rss, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&rsstest1.Rss{})
    var rsss []rsstest1.Rss
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&rsss).Error
	return  rsss, total, err
}
