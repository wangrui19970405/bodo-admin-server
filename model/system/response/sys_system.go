package response

import "github.com/wangrui19970405/wu-shi-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
