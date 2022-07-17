package utils

var adminId = [...]int64{1}

func IsAdmin(userId int64) (b bool) {
	b = false
	for _, v := range adminId {
		if v == userId {
			b = true
			return
		}
	}
	return
}
