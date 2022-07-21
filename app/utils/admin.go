package utils

var adminId = map[int64]struct{}{1: {}}

func IsAdmin(userId int64) (b bool) {
	_, b = adminId[userId]
	return
}
