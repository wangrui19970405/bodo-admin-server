package request

import (
	"github.com/wangrui19970405/wu-shi-admin/server/model/common/request"
	"github.com/wangrui19970405/wu-shi-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
