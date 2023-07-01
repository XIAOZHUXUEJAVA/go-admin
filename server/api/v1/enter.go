package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/rsstest1"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/usertest1"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/usertest2"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	PkgTestApiGroup   pkgTest.ApiGroup
	Usertest1ApiGroup usertest1.ApiGroup
	Usertest2ApiGroup usertest2.ApiGroup
	Rsstest1ApiGroup  rsstest1.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
