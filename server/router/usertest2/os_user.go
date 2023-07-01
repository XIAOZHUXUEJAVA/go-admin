package usertest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OSUserRouter struct {
}

// InitOSUserRouter 初始化 OSUser 路由信息
func (s *OSUserRouter) InitOSUserRouter(Router *gin.RouterGroup) {
	OSuserRouter := Router.Group("OSuser").Use(middleware.OperationRecord())
	OSuserRouterWithoutRecord := Router.Group("OSuser")
	var OSuserApi = v1.ApiGroupApp.Usertest2ApiGroup.OSUserApi
	{
		OSuserRouter.POST("createOSUser", OSuserApi.CreateOSUser)   // 新建OSUser
		OSuserRouter.DELETE("deleteOSUser", OSuserApi.DeleteOSUser) // 删除OSUser
		OSuserRouter.DELETE("deleteOSUserByIds", OSuserApi.DeleteOSUserByIds) // 批量删除OSUser
		OSuserRouter.PUT("updateOSUser", OSuserApi.UpdateOSUser)    // 更新OSUser
	}
	{
		OSuserRouterWithoutRecord.GET("findOSUser", OSuserApi.FindOSUser)        // 根据ID获取OSUser
		OSuserRouterWithoutRecord.GET("getOSUserList", OSuserApi.GetOSUserList)  // 获取OSUser列表
	}
}
