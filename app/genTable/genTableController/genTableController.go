package genTableController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGenTableController)
