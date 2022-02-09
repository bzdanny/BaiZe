package baizeContext

import (
	"baize/app/constant/constants"
	"baize/app/system/models/loginModels"
)

func (bzc *BaiZeContext) GetCurrentLoginUser() (loginUser *loginModels.LoginUser) {
	loginUserKey, _ := bzc.Get(constants.LoginUserKey)
	if loginUserKey != nil {
		loginUser = loginUserKey.(*loginModels.LoginUser)
	}
	return
}
func (bzc *BaiZeContext) GetCurrentUser() (user *loginModels.User) {
	user = bzc.GetCurrentLoginUser().User
	if user == nil {
		return nil
	}
	return

}
func (bzc *BaiZeContext) GetCurrentUserName() string {
	user := bzc.GetCurrentUser()
	if user == nil {
		return ""
	}
	return user.UserName
}
func (bzc *BaiZeContext) GetCurrentUserId() int64 {
	user := bzc.GetCurrentUser()
	if user == nil {
		return 0
	}
	return user.UserId
}
