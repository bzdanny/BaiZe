package main

import (
	"baize/app/common/mysql"
	"baize/app/common/redis"
	"baize/app/routes"
	"baize/app/setting"
	"baize/app/utils"
	"baize/app/utils/logger"
	"fmt"
	"os"
)

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
