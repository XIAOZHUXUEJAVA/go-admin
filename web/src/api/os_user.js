import service from '@/utils/request'

// @Tags OSUser
// @Summary 创建OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OSUser true "创建OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /OSuser/createOSUser [post]
export const createOSUser = (data) => {
  return service({
    url: '/OSuser/createOSUser',
    method: 'post',
    data
  })
}

// @Tags OSUser
// @Summary 删除OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OSUser true "删除OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /OSuser/deleteOSUser [delete]
export const deleteOSUser = (data) => {
  return service({
    url: '/OSuser/deleteOSUser',
    method: 'delete',
    data
  })
}

// @Tags OSUser
// @Summary 删除OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /OSuser/deleteOSUser [delete]
export const deleteOSUserByIds = (data) => {
  return service({
    url: '/OSuser/deleteOSUserByIds',
    method: 'delete',
    data
  })
}

// @Tags OSUser
// @Summary 更新OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OSUser true "更新OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /OSuser/updateOSUser [put]
export const updateOSUser = (data) => {
  return service({
    url: '/OSuser/updateOSUser',
    method: 'put',
    data
  })
}

// @Tags OSUser
// @Summary 用id查询OSUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OSUser true "用id查询OSUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /OSuser/findOSUser [get]
export const findOSUser = (params) => {
  return service({
    url: '/OSuser/findOSUser',
    method: 'get',
    params
  })
}

// @Tags OSUser
// @Summary 分页获取OSUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OSUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /OSuser/getOSUserList [get]
export const getOSUserList = (params) => {
  return service({
    url: '/OSuser/getOSUserList',
    method: 'get',
    params
  })
}
