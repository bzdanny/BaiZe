package monitorService

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
)

type IUserOnlineService interface {
	SelectUserOnlineList(ipaddr, userName string) (list []*monitorModels.SysUserOnline, total *int64)
	ForceLogout(tokenId string)
}
