package main

import (
	"go.uber.org/zap"

	"github.com/wangrui19970405/wu-shi-admin/server/core"
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.BODO_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.BODO_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.BODO_LOG)
	global.BODO_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.BODO_DB != nil {
		initialize.RegisterTables(global.BODO_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.BODO_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
