package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type FoodApi struct {
}

var foodService = service.ServiceGroupApp.PkgTestServiceGroup.FoodService


// CreateFood 创建Food
// @Tags Food
// @Summary 创建Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Food true "创建Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /food/createFood [post]
func (foodApi *FoodApi) CreateFood(c *gin.Context) {
	var food pkgTest.Food
	err := c.ShouldBindJSON(&food)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := foodService.CreateFood(&food); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFood 删除Food
// @Tags Food
// @Summary 删除Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Food true "删除Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /food/deleteFood [delete]
func (foodApi *FoodApi) DeleteFood(c *gin.Context) {
	var food pkgTest.Food
	err := c.ShouldBindJSON(&food)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := foodService.DeleteFood(food); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFoodByIds 批量删除Food
// @Tags Food
// @Summary 批量删除Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /food/deleteFoodByIds [delete]
func (foodApi *FoodApi) DeleteFoodByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := foodService.DeleteFoodByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFood 更新Food
// @Tags Food
// @Summary 更新Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.Food true "更新Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /food/updateFood [put]
func (foodApi *FoodApi) UpdateFood(c *gin.Context) {
	var food pkgTest.Food
	err := c.ShouldBindJSON(&food)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := foodService.UpdateFood(food); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFood 用id查询Food
// @Tags Food
// @Summary 用id查询Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.Food true "用id查询Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /food/findFood [get]
func (foodApi *FoodApi) FindFood(c *gin.Context) {
	var food pkgTest.Food
	err := c.ShouldBindQuery(&food)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refood, err := foodService.GetFood(food.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refood": refood}, c)
	}
}

// GetFoodList 分页获取Food列表
// @Tags Food
// @Summary 分页获取Food列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.FoodSearch true "分页获取Food列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /food/getFoodList [get]
func (foodApi *FoodApi) GetFoodList(c *gin.Context) {
	var pageInfo pkgTestReq.FoodSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := foodService.GetFoodInfoList(pageInfo); err != nil {
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
