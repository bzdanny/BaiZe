//go:build wireinject
// +build wireinject

package main

import (
	"baize/app/routes"
	"baize/app/setting"
	"baize/baize/datasource"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func wireApp(*setting.Datasource) (*gin.Engine, func(), error) {
	panic(wire.Build(
		datasource.ProviderSet,
		routes.ProviderSet,
		newApp))
}
