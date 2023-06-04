package system

import (
	"github.com/wangrui19970405/wu-shi-admin/server/global"
)

type JwtBlacklist struct {
	global.BODO_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
