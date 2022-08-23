package baizeContext

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/constants"
)

func (bzc *BaiZeContext) GetCurrentUser() (loginUser *systemModels.LoginUser) {
	loginUserKey, _ := bzc.Get(constants.LoginUserKey)
	if loginUserKey != nil {
		loginUser = loginUserKey.(*systemModels.LoginUser)
	}
	return
}
func (bzc *BaiZeContext) GetUser() (user *systemModels.User) {
	user = bzc.GetCurrentUser().User
	if user == nil {
		return nil
	}
	return

}
func (bzc *BaiZeContext) GetUserName() string {
	user := bzc.GetUser()
	if user == nil {
		return ""
	}
	return user.UserName
}
func (bzc *BaiZeContext) GetUserId() int64 {
	user := bzc.GetUser()
	if user == nil {
		return 0
	}
	return user.UserId
}
