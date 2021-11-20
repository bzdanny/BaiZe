package routes

import (
	"baize/app/common/middlewares"
	"baize/app/constant/constants"
	"baize/app/routes/genTableRoutes"
	"baize/app/routes/monitor"
	"baize/app/routes/systemRouter"
	"baize/app/setting"
	"baize/app/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	if setting.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	group := r.Group("")
	group.Static(constants.ResourcePrefix, setting.Conf.Profile)
	//不做鉴权的
	{
		systemRouter.InitLoginRouter(group) //获取登录信息
	}
	//做鉴权的
	group.Use(middlewares.JWTAuthMiddleware())
	{

		systemRouter.InitGetUser(group)           //获取登录信息
		systemRouter.InitSysProfileRouter(group)  //个人信息
		systemRouter.InitSysUserRouter(group)     //用户相关
		systemRouter.InitSysDeptRouter(group)     //部门相关
		systemRouter.InitSysDictDataRouter(group) //数据字典信息
		systemRouter.InitSysRoleRouter(group)     //角色相关
		systemRouter.InitSysMenuRouter(group)     //菜单相关
		systemRouter.InitSysConfigRouter(group)   //参数配置
		systemRouter.InitSysDictTypeRouter(group) //数据字典属性
		systemRouter.InitSysPostRouter(group)     //岗位属性
		monitor.InitSysUserOnlineRouter(group)    //在线用户监控
		monitor.InitSysLogininforRouter(group)    //登录用户日志
		monitor.InitSysOperLogRouter(group)       //操作日志
		monitor.InitServerRouter(group)           //服务监控
		genTableRoutes.InitGenTableRouter(group)  //代码生成

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r

}
