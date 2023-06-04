package v1

import (
	"github.com/wangrui19970405/wu-shi-admin/server/api/v1/example"
	"github.com/wangrui19970405/wu-shi-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
