package usertest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/usertest2"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    usertest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/usertest2/request"
)

type OSUserService struct {
}

// CreateOSUser 创建OSUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (OSuserService *OSUserService) CreateOSUser(OSuser *usertest2.OSUser) (err error) {
	err = global.GVA_DB.Create(OSuser).Error
	return err
}

// DeleteOSUser 删除OSUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (OSuserService *OSUserService)DeleteOSUser(OSuser usertest2.OSUser) (err error) {
	err = global.GVA_DB.Delete(&OSuser).Error
	return err
}

// DeleteOSUserByIds 批量删除OSUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (OSuserService *OSUserService)DeleteOSUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]usertest2.OSUser{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOSUser 更新OSUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (OSuserService *OSUserService)UpdateOSUser(OSuser usertest2.OSUser) (err error) {
	err = global.GVA_DB.Save(&OSuser).Error
	return err
}

// GetOSUser 根据id获取OSUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (OSuserService *OSUserService)GetOSUser(id uint) (OSuser usertest2.OSUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&OSuser).Error
	return
}

// GetOSUserInfoList 分页获取OSUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (OSuserService *OSUserService)GetOSUserInfoList(info usertest2Req.OSUserSearch) (list []usertest2.OSUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&usertest2.OSUser{})
    var OSusers []usertest2.OSUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&OSusers).Error
	return  OSusers, total, err
}
