package middleware

import (
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/model/common/response"
	"github.com/wangrui19970405/wu-shi-admin/server/service"
	"github.com/wangrui19970405/wu-shi-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.WUSHI_CONFIG.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c)
			//获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
