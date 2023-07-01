import service from '@/utils/request'

// @Tags Rss
// @Summary 创建Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rss true "创建Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rss/createRss [post]
export const createRss = (data) => {
  return service({
    url: '/rss/createRss',
    method: 'post',
    data
  })
}

// @Tags Rss
// @Summary 删除Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rss true "删除Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rss/deleteRss [delete]
export const deleteRss = (data) => {
  return service({
    url: '/rss/deleteRss',
    method: 'delete',
    data
  })
}

// @Tags Rss
// @Summary 删除Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rss/deleteRss [delete]
export const deleteRssByIds = (data) => {
  return service({
    url: '/rss/deleteRssByIds',
    method: 'delete',
    data
  })
}

// @Tags Rss
// @Summary 更新Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rss true "更新Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rss/updateRss [put]
export const updateRss = (data) => {
  return service({
    url: '/rss/updateRss',
    method: 'put',
    data
  })
}

// @Tags Rss
// @Summary 用id查询Rss
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Rss true "用id查询Rss"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rss/findRss [get]
export const findRss = (params) => {
  return service({
    url: '/rss/findRss',
    method: 'get',
    params
  })
}

// @Tags Rss
// @Summary 分页获取Rss列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Rss列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rss/getRssList [get]
export const getRssList = (params) => {
  return service({
    url: '/rss/getRssList',
    method: 'get',
    params
  })
}
