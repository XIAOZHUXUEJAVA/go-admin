package rsstest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RssRouter struct {
}

// InitRssRouter 初始化 Rss 路由信息
func (s *RssRouter) InitRssRouter(Router *gin.RouterGroup) {
	rssRouter := Router.Group("rss").Use(middleware.OperationRecord())
	rssRouterWithoutRecord := Router.Group("rss")
	var rssApi = v1.ApiGroupApp.Rsstest1ApiGroup.RssApi
	{
		rssRouter.POST("createRss", rssApi.CreateRss)   // 新建Rss
		rssRouter.DELETE("deleteRss", rssApi.DeleteRss) // 删除Rss
		rssRouter.DELETE("deleteRssByIds", rssApi.DeleteRssByIds) // 批量删除Rss
		rssRouter.PUT("updateRss", rssApi.UpdateRss)    // 更新Rss
	}
	{
		rssRouterWithoutRecord.GET("findRss", rssApi.FindRss)        // 根据ID获取Rss
		rssRouterWithoutRecord.GET("getRssList", rssApi.GetRssList)  // 获取Rss列表
	}
}
