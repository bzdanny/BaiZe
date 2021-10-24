package systemService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/utils/bCryptPasswordEncoder"
)

func MatchesPassword(rawPassword string, userId int64) bool {

	return bCryptPasswordEncoder.CheckPasswordHash(rawPassword, systemDao.SelectPasswordByUserId(userId))
}
