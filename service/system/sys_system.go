package system

import (
	"github.com/wangrui19970405/wu-shi-admin/server/config"
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system"
	"github.com/wangrui19970405/wu-shi-admin/server/utils"
	"go.uber.org/zap"
)

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf config.Server, err error

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return global.WUSHI_CONFIG, nil
}

// @description   set system config,
//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.WUSHI_VP.Set(k, v)
	}
	err = global.WUSHI_VP.WriteConfig()
	return err
}

//@author: [wangrui19970405](https://github.com/wangrui19970405)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.WUSHI_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.WUSHI_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.WUSHI_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
