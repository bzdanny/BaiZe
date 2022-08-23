package genTableServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(GetGenTabletColumnService, GetGenTabletService)
