package main

import (
	"baize/app/routes"
	"baize/app/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	if len(os.Args) <= 1 {
		setting.Init("../config/config.yaml")
	} else {
		setting.Init(os.Args[1])
	}

}

func newApp() *gin.Engine {

	return routes.RegisterServer()
}

// staging.knithq.com/knit
// dev-knit.ibaize.vip/knit
// localhost:9000/knit

// @title knit
// @version 1.0.x
// @description knit接口文档

// @contact.name danny
// @contact.email zdyang@knitpeople.com

// @host localhost:9000/knit
//// @host dev-knit.ibaize.vip/knit

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	app.Run(fmt.Sprintf(":%d", setting.Conf.Port))

}
