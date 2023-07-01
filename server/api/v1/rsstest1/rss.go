package rsstest1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/rsstest1"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    rsstest1Req "github.com/flipped-aurora/gin-vue-admin/server/model/rsstest1/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type RssApi struct {
}

var rssService = service.ServiceGroupApp.Rsstest1ServiceGroup.RssService


// CreateRss 创建Rss
// @Tags Rss
// @Summary 创建Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rsstest1.Rss true "创建Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rss/createRss [post]
func (rssApi *RssApi) CreateRss(c *gin.Context) {
	var rss rsstest1.Rss
	err := c.ShouldBindJSON(&rss)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rssService.CreateRss(&rss); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRss 删除Rss
// @Tags Rss
// @Summary 删除Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rsstest1.Rss true "删除Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rss/deleteRss [delete]
func (rssApi *RssApi) DeleteRss(c *gin.Context) {
	var rss rsstest1.Rss
	err := c.ShouldBindJSON(&rss)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rssService.DeleteRss(rss); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRssByIds 批量删除Rss
// @Tags Rss
// @Summary 批量删除Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rss/deleteRssByIds [delete]
func (rssApi *RssApi) DeleteRssByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rssService.DeleteRssByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRss 更新Rss
// @Tags Rss
// @Summary 更新Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body rsstest1.Rss true "更新Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rss/updateRss [put]
func (rssApi *RssApi) UpdateRss(c *gin.Context) {
	var rss rsstest1.Rss
	err := c.ShouldBindJSON(&rss)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rssService.UpdateRss(rss); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRss 用id查询Rss
// @Tags Rss
// @Summary 用id查询Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rsstest1.Rss true "用id查询Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rss/findRss [get]
func (rssApi *RssApi) FindRss(c *gin.Context) {
	var rss rsstest1.Rss
	err := c.ShouldBindQuery(&rss)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerss, err := rssService.GetRss(rss.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerss": rerss}, c)
	}
}

// GetRssList 分页获取Rss列表
// @Tags Rss
// @Summary 分页获取Rss列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rsstest1Req.RssSearch true "分页获取Rss列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rss/getRssList [get]
func (rssApi *RssApi) GetRssList(c *gin.Context) {
	var pageInfo rsstest1Req.RssSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := rssService.GetRssInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
