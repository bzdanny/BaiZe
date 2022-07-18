package baizeContext

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
)

func (bzc *BaiZeContext) GetCurrentLoginUser() (loginUser *systemModels.LoginUser) {
	loginUserKey, _ := bzc.Get(constants.LoginUserKey)
	if loginUserKey != nil {
		loginUser = loginUserKey.(*systemModels.LoginUser)
	}
	return
}
func (bzc *BaiZeContext) GetCurrentUser() (user *systemModels.User) {
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
