package monitorServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(GetLogininforService, NewOperLogServiceService, NewUserOnlineService)
