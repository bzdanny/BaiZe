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
	if err := setting.Init(filePath); err != nil {
		fmt.Printf("init setting 失败，错误:%v\n", err)
	}
	//2.初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger 失败，错误:%v\n", err)
	}
	//3.初始化MySQL
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql 失败，错误:%v\n", err)
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	//4.初始化Redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis 失败，错误:%v\n", err)
	}
	defer redis.Close()

	if err := utils.Init(); err != nil {
		fmt.Printf("init utils 失败，错误:%v\n", err)
	}
	//5.注册路由启动服务
	r := routes.Init(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
	fmt.Println("启动成功")
}
