package loginService

import (
	"baize/app/monitor/monitorModels"
	"baize/app/system/models/loginModels"
)

type ILoginService interface {
	Login(user *loginModels.User, l *monitorModels.Logininfor) *string
	RecordLoginInfo(loginUser *monitorModels.Logininfor)
}
