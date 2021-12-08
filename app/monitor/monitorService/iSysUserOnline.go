package monitorService

import (
	"baize/app/monitor/monitorModels"
)

type ItUserOnlineService interface {
	SelectUserOnlineList(ipaddr, userName string) (list []*monitorModels.SysUserOnline, total *int64)
	ForceLogout(tokenId string)
}
