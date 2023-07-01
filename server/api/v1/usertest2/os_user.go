package usertest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/usertest2"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    usertest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/usertest2/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OSUserApi struct {
}

var OSuserService = service.ServiceGroupApp.Usertest2ServiceGroup.OSUserService


// CreateOSUser 创建OSUser
// @Tags OSUser
// @Summary 创建OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body usertest2.OSUser true "创建OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /OSuser/createOSUser [post]
func (OSuserApi *OSUserApi) CreateOSUser(c *gin.Context) {
	var OSuser usertest2.OSUser
	err := c.ShouldBindJSON(&OSuser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := OSuserService.CreateOSUser(&OSuser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOSUser 删除OSUser
// @Tags OSUser
// @Summary 删除OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body usertest2.OSUser true "删除OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /OSuser/deleteOSUser [delete]
func (OSuserApi *OSUserApi) DeleteOSUser(c *gin.Context) {
	var OSuser usertest2.OSUser
	err := c.ShouldBindJSON(&OSuser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := OSuserService.DeleteOSUser(OSuser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOSUserByIds 批量删除OSUser
// @Tags OSUser
// @Summary 批量删除OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /OSuser/deleteOSUserByIds [delete]
func (OSuserApi *OSUserApi) DeleteOSUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := OSuserService.DeleteOSUserByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOSUser 更新OSUser
// @Tags OSUser
// @Summary 更新OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body usertest2.OSUser true "更新OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /OSuser/updateOSUser [put]
func (OSuserApi *OSUserApi) UpdateOSUser(c *gin.Context) {
	var OSuser usertest2.OSUser
	err := c.ShouldBindJSON(&OSuser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := OSuserService.UpdateOSUser(OSuser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOSUser 用id查询OSUser
// @Tags OSUser
// @Summary 用id查询OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query usertest2.OSUser true "用id查询OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /OSuser/findOSUser [get]
func (OSuserApi *OSUserApi) FindOSUser(c *gin.Context) {
	var OSuser usertest2.OSUser
	err := c.ShouldBindQuery(&OSuser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reOSuser, err := OSuserService.GetOSUser(OSuser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reOSuser": reOSuser}, c)
	}
}

// GetOSUserList 分页获取OSUser列表
// @Tags OSUser
// @Summary 分页获取OSUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query usertest2Req.OSUserSearch true "分页获取OSUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /OSuser/getOSUserList [get]
func (OSuserApi *OSUserApi) GetOSUserList(c *gin.Context) {
	var pageInfo usertest2Req.OSUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := OSuserService.GetOSUserInfoList(pageInfo); err != nil {
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
