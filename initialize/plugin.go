package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/middleware"
	"github.com/wangrui19970405/wu-shi-admin/server/plugin/email"
	"github.com/wangrui19970405/wu-shi-admin/server/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权分包模块安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("带鉴权分包模块安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.WUSHI_CONFIG.Email.To,
		global.WUSHI_CONFIG.Email.From,
		global.WUSHI_CONFIG.Email.Host,
		global.WUSHI_CONFIG.Email.Secret,
		global.WUSHI_CONFIG.Email.Nickname,
		global.WUSHI_CONFIG.Email.Port,
		global.WUSHI_CONFIG.Email.IsSSL,
	))
}
