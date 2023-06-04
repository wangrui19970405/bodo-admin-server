package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/wangrui19970405/wu-shi-admin/server/config"
	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"github.com/wangrui19970405/wu-shi-admin/server/utils"
)

func Timer() {
	if global.WUSHI_CONFIG.Timer.Start {
		for i := range global.WUSHI_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.WUSHI_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.WUSHI_Timer.AddTaskByFunc("ClearDB", global.WUSHI_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.WUSHI_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.WUSHI_CONFIG.Timer.Detail[i])
		}
	}
}
