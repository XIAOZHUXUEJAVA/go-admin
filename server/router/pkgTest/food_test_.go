package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FoodRouter struct {
}

// InitFoodRouter 初始化 Food 路由信息
func (s *FoodRouter) InitFoodRouter(Router *gin.RouterGroup) {
	foodRouter := Router.Group("food").Use(middleware.OperationRecord())
	foodRouterWithoutRecord := Router.Group("food")
	var foodApi = v1.ApiGroupApp.PkgTestApiGroup.FoodApi
	{
		foodRouter.POST("createFood", foodApi.CreateFood)   // 新建Food
		foodRouter.DELETE("deleteFood", foodApi.DeleteFood) // 删除Food
		foodRouter.DELETE("deleteFoodByIds", foodApi.DeleteFoodByIds) // 批量删除Food
		foodRouter.PUT("updateFood", foodApi.UpdateFood)    // 更新Food
	}
	{
		foodRouterWithoutRecord.GET("findFood", foodApi.FindFood)        // 根据ID获取Food
		foodRouterWithoutRecord.GET("getFoodList", foodApi.GetFoodList)  // 获取Food列表
	}
}
