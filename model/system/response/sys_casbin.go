package response

import (
	"github.com/wangrui19970405/wu-shi-admin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
