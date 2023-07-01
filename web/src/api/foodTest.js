import service from '@/utils/request'

// @Tags Food
// @Summary 创建Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Food true "创建Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /food/createFood [post]
export const createFood = (data) => {
  return service({
    url: '/food/createFood',
    method: 'post',
    data
  })
}

// @Tags Food
// @Summary 删除Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Food true "删除Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /food/deleteFood [delete]
export const deleteFood = (data) => {
  return service({
    url: '/food/deleteFood',
    method: 'delete',
    data
  })
}

// @Tags Food
// @Summary 删除Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /food/deleteFood [delete]
export const deleteFoodByIds = (data) => {
  return service({
    url: '/food/deleteFoodByIds',
    method: 'delete',
    data
  })
}

// @Tags Food
// @Summary 更新Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Food true "更新Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /food/updateFood [put]
export const updateFood = (data) => {
  return service({
    url: '/food/updateFood',
    method: 'put',
    data
  })
}

// @Tags Food
// @Summary 用id查询Food
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Food true "用id查询Food"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /food/findFood [get]
export const findFood = (params) => {
  return service({
    url: '/food/findFood',
    method: 'get',
    params
  })
}

// @Tags Food
// @Summary 分页获取Food列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Food列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /food/getFoodList [get]
export const getFoodList = (params) => {
  return service({
    url: '/food/getFoodList',
    method: 'get',
    params
  })
}
