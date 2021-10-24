package systemService

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/system/models/systemModels"
)

func SelectUserOnlineList(ipaddr, userName string) (list []*systemModels.SysUserOnline, total *int64) {
	keys := redis.Keys(constants.LoginTokenKey + "*")
	list = make([]*systemModels.SysUserOnline, 0, len(keys))
	for _, key := range keys {
		user, err := redis.GetCacheLoginUser(key)
		if err == nil && (ipaddr == "" || ipaddr == user.IpAddr) && (userName == "" || userName == user.User.UserName) {
			list = append(list, systemModels.GetSysUserOnlineByUser(user))
		}
	}
	i := int64(len(list))
	total = &i
	return
}

func ForceLogout(tokenId string) {
	redis.Delete(constants.LoginTokenKey + tokenId)
}
