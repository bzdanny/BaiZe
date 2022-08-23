package genTableDaoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGenTableColumnDao, NewGenTableDao)
