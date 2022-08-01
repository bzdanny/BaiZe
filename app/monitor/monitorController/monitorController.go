package monitorController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMonitorController, NewInfoServerController, NewLogininforController, NewOperLogController, NewUserOnlineController)

type MonitorController struct {
	Info       *InfoServerController
	Logininfor *LogininforController
	OperLog    *OperLogController
	UserOnline *UserOnlineController
}

func NewMonitorController(
	Info *InfoServerController,
	Logininfor *LogininforController,
	OperLog *OperLogController,
	UserOnline *UserOnlineController,
) *MonitorController {
	return &MonitorController{
		Info:       Info,
		Logininfor: Logininfor,
		OperLog:    OperLog,
		UserOnline: UserOnline,
	}
}
