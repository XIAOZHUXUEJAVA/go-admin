package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/rsstest1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type RssSearch struct{
    rsstest1.Rss
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
