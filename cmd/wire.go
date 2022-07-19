//go:build wireinject
// +build wireinject

package main

import (
	"baize/app/routes"
	"baize/app/setting"
	"baize/baize/datasource"
	"github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
)

func wireApp(*setting.Datasource) (*gin.Engine, func(), error) {
	panic(wire.Build(
		datasource.ProviderSet,
		systemDaoImpl.ProviderSet,
		systemServiceImpl.ProviderSet,
		routes.ProviderSet,
		newApp))
}
