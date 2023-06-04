package system

import (
	"github.com/wangrui19970405/wu-shi-admin/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
