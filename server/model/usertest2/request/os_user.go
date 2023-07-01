package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/usertest2"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OSUserSearch struct{
    usertest2.OSUser
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
