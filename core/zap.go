package core

import (
	"fmt"
	"github.com/wangrui19970405/wu-shi-admin/server/core/internal"
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [wangrui19970405](https://github.com/wangrui19970405)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.WUSHI_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.WUSHI_CONFIG.Zap.Director)
		_ = os.Mkdir(global.WUSHI_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.WUSHI_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
