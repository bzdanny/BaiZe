package monitorService

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/monitor/monitorModels"
)

func SelectUserOnlineList(ipaddr, userName string) (list []*monitorModels.SysUserOnline, total *int64) {
	keys := redis.Keys(constants.LoginTokenKey + "*")
	list = make([]*monitorModels.SysUserOnline, 0, len(keys))
	for _, key := range keys {
		user, err := redis.GetCacheLoginUser(key)
		if err == nil && (ipaddr == "" || ipaddr == user.IpAddr) && (userName == "" || userName == user.User.UserName) {
			list = append(list, monitorModels.GetSysUserOnlineByUser(user))
		}
	}
	i := int64(len(list))
	total = &i
	return
}

func ForceLogout(tokenId string) {
	redis.Delete(constants.LoginTokenKey + tokenId)
}
