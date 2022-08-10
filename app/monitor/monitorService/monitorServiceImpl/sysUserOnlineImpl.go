package monitorServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/constants"
	"github.com/bzdanny/BaiZe/baize/utils/redisUtils"
)

type UserOnlineService struct {
}

func NewUserOnlineService() *UserOnlineService {
	return new(UserOnlineService)
}

func (userOnlineService *UserOnlineService) SelectUserOnlineList(ipaddr, userName string) (list []*monitorModels.SysUserOnline, total *int64) {
	keys := redisUtils.Keys(constants.LoginTokenKey + "*")
	list = make([]*monitorModels.SysUserOnline, 0, len(keys))
	loginUser := new(systemModels.LoginUser)
	for _, key := range keys {
		user, err := redisUtils.GetStruct(key, loginUser)
		if err == nil && (ipaddr == "" || ipaddr == user.IpAddr) && (userName == "" || userName == user.User.UserName) {
			list = append(list, monitorModels.GetSysUserOnlineByUser(user))
		}
	}
	i := int64(len(list))
	total = &i
	return
}

func (userOnlineService *UserOnlineService) ForceLogout(tokenId string) {
	redisUtils.Delete(constants.LoginTokenKey + tokenId)
}
