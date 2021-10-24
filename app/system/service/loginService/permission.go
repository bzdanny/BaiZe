package loginService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/utils/admin"
)

func GetMenuPermission(userId int64) []string {
	perms := make([]string, 0, 1)
	if admin.IsAdmin(userId) {
		perms = append(perms, "*:*:*")
	} else {
		mysqlPerms := systemDao.SelectMenuPermsByUserId(userId)

		for _, perm := range mysqlPerms {
			if len(perm) != 0 {
				perms = append(perms, perm)
			}
		}
	}
	return perms
}
