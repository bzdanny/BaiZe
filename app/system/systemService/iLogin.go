package systemService

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/monitor/monitorModels"
)

type ILoginService interface {
	Login(user *systemModels.User, l *monitorModels.Logininfor) *string
	RecordLoginInfo(loginUser *monitorModels.Logininfor)
}
