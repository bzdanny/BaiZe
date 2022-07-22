package main

import (
	"fmt"
	"github.com/bzdanny/BaiZe/app/routes"
	"github.com/bzdanny/BaiZe/app/setting"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	if len(os.Args) <= 1 {
		setting.Init("./config/config.yaml")
	} else {
		setting.Init(os.Args[1])
	}

}

func newApp(r *routes.Router) *gin.Engine {

	return routes.RegisterServer(r)
}

// localhost:9000

// @title baize
// @version 2.0.x
// @description baize接口文档

// @contact.name danny
// @contact.email zhao_402295440@126.com

// @host localhost:9000

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	app, cleanup, err := wireApp(setting.Conf.Datasource)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	app.Run(fmt.Sprintf(":%d", setting.Conf.Port))

}
