//go:build wireinject
// +build wireinject

package main

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorController"
	"github.com/bzdanny/BaiZe/app/monitor/monitorDao/monitorDaoImpl"
	"github.com/bzdanny/BaiZe/app/monitor/monitorService/monitorServiceImpl"
	"github.com/bzdanny/BaiZe/app/routes"
	"github.com/bzdanny/BaiZe/app/setting"
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/app/system/systemDao/systemDaoImpl"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func wireApp(*setting.Datasource) (*gin.Engine, func(), error) {
	panic(wire.Build(
		datasource.ProviderSet,
		monitorController.ProviderSet,
		monitorServiceImpl.ProviderSet,
		monitorDaoImpl.ProviderSet,
		systemDaoImpl.ProviderSet,
		systemServiceImpl.ProviderSet,
		systemController.ProviderSet,
		routes.ProviderSet,
		newApp))
}
