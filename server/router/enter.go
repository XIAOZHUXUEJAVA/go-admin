package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/router/rsstest1"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/usertest1"
	"github.com/flipped-aurora/gin-vue-admin/server/router/usertest2"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	PkgTest   pkgTest.RouterGroup
	Usertest1 usertest1.RouterGroup
	Usertest2 usertest2.RouterGroup
	Rsstest1  rsstest1.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
