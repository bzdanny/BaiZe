package systemService

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
)

type ILoginService interface {
	Login(user *systemModels.User, l *monitorModels.Logininfor) *string
	RecordLoginInfo(loginUser *monitorModels.Logininfor)
	GenerateCode() (m *systemModels.CaptchaVo)
	VerityCaptcha(id, base64 string) bool
}
