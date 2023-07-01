package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/service/rsstest1"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/usertest1"
	"github.com/flipped-aurora/gin-vue-admin/server/service/usertest2"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	PkgTestServiceGroup   pkgTest.ServiceGroup
	Usertest1ServiceGroup usertest1.ServiceGroup
	Usertest2ServiceGroup usertest2.ServiceGroup
	Rsstest1ServiceGroup  rsstest1.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
