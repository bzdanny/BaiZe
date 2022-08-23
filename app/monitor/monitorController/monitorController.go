package monitorController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMonitorController, NewInfoServerController, NewLogininforController, NewUserOnlineController)

type MonitorController struct {
	Info       *InfoServerController
	Logininfor *LogininforController
	UserOnline *UserOnlineController
}

func NewMonitorController(
	Info *InfoServerController,
	Logininfor *LogininforController,
	UserOnline *UserOnlineController,
) *MonitorController {
	return &MonitorController{
		Info:       Info,
		Logininfor: Logininfor,
		UserOnline: UserOnline,
	}
}
