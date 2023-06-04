package service

import (
	"github.com/wangrui19970405/wu-shi-admin/server/service/example"
	"github.com/wangrui19970405/wu-shi-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
