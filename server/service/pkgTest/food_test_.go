package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
)

type FoodService struct {
}

// CreateFood 创建Food记录
// Author [piexlmax](https://github.com/piexlmax)
func (foodService *FoodService) CreateFood(food *pkgTest.Food) (err error) {
	err = global.GVA_DB.Create(food).Error
	return err
}

// DeleteFood 删除Food记录
// Author [piexlmax](https://github.com/piexlmax)
func (foodService *FoodService)DeleteFood(food pkgTest.Food) (err error) {
	err = global.GVA_DB.Delete(&food).Error
	return err
}

// DeleteFoodByIds 批量删除Food记录
// Author [piexlmax](https://github.com/piexlmax)
func (foodService *FoodService)DeleteFoodByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest.Food{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFood 更新Food记录
// Author [piexlmax](https://github.com/piexlmax)
func (foodService *FoodService)UpdateFood(food pkgTest.Food) (err error) {
	err = global.GVA_DB.Save(&food).Error
	return err
}

// GetFood 根据id获取Food记录
// Author [piexlmax](https://github.com/piexlmax)
func (foodService *FoodService)GetFood(id uint) (food pkgTest.Food, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&food).Error
	return
}

// GetFoodInfoList 分页获取Food记录
// Author [piexlmax](https://github.com/piexlmax)
func (foodService *FoodService)GetFoodInfoList(info pkgTestReq.FoodSearch) (list []pkgTest.Food, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.Food{})
    var foods []pkgTest.Food
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&foods).Error
	return  foods, total, err
}
