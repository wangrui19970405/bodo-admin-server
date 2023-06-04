package router

import (
	"github.com/wangrui19970405/wu-shi-admin/server/router/example"
	"github.com/wangrui19970405/wu-shi-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
