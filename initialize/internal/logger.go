package internal

import (
	"fmt"

	"github.com/wangrui19970405/wu-shi-admin/server/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [wangrui19970405](https://github.com/wangrui19970405)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [wangrui19970405](https://github.com/wangrui19970405)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.WUSHI_CONFIG.System.DbType {
	case "mysql":
		logZap = global.WUSHI_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.WUSHI_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.WUSHI_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
