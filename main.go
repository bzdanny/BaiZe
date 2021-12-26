package main

import (
	"baize/app/common/commonModels"
	"baize/app/common/mysql"
	"baize/app/common/redis"
	"baize/app/routes"
	"baize/app/setting"
	"baize/app/utils"
	"baize/app/utils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)


// @title 白泽
// @version 1.0.x
// @description 白泽接口文档
// @termsOfService https://www.ibaize.vip

// @contact.name danny
// @contact.url https://www.ibaize.vip
// @contact.email zhao_402295440@126.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080



func main() {

	var filePath string
	if len(os.Args) <= 1 {
		filePath = "./config.yaml"
	} else {
		filePath = os.Args[1]
	}
	//1.加载配置
	setting.Init(filePath)
	//2.初始化日志
	logger.Init()
	//3.初始化MySQL
	mysql.Init()
	defer mysql.Close() // 程序退出关闭数据库连接
	//4.初始化Redis
	redis.Init()
	defer redis.Close()
	//5.初始化Utils
	utils.Init()
	//6.注册路由启动服务
	r := routes.Init()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		panic(err)
	}
}


// @Summary 添加用户
// @accept multipart/form-data
// @Produce json
// @Param email formData string true "用户邮箱地址"
// @Param username formData string true "用户名"
// @Param password formData string true "用户密码"
// @Success 200 {string} json "{"code":200,"data":[],"msg":{"title":"ok"}}"
// @Router /api/v1/users [post]

func DemoUserList(c *gin.Context) {
	c.JSON(http.StatusOK, commonModels.Success())

}