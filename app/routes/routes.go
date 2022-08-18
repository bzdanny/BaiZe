package routes

import (
	"github.com/bzdanny/BaiZe/app/genTable/genTableController"
	"github.com/bzdanny/BaiZe/app/monitor/monitorController"
	"github.com/bzdanny/BaiZe/app/routes/monitorRoutes"
	"github.com/bzdanny/BaiZe/app/routes/systemRoutes"
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/IOFile"
	"github.com/bzdanny/BaiZe/baize/constants"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/bzdanny/BaiZe/baize/setting"
	"github.com/bzdanny/BaiZe/baize/utils/logger"
	"github.com/google/wire"
	"github.com/swaggo/files"

	"fmt"
	"net/http"
	"strings"

	"github.com/bzdanny/BaiZe/app/docs"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

var ProviderSet = wire.NewSet(NewRouter)

type Router struct {
	Sys      *systemController.SystemController
	Monitor  *monitorController.MonitorController
	GenTable *genTableController.GenTableController
}

func NewRouter(sys *systemController.SystemController,
	monitor *monitorController.MonitorController) *Router {
	return &Router{
		Sys:     sys,
		Monitor: monitor,
	}
}

func RegisterServer(router *Router) *gin.Engine {

	if setting.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(Cors())
	group := r.Group("")

	host := setting.Conf.Host
	docs.SwaggerInfo.Host = host[strings.Index(host, "//")+2:]
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	group.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//如果是本地存储则开启
	if !IOFile.FileType.Contains(setting.Conf.UploadFile.Type) {
		path := setting.Conf.UploadFile.Localhost.PublicResourcePrefix
		if path == "" {
			path = constants.DefaultPublicPath
		}
		group.Static(constants.ResourcePrefix, path)
	}

	//不做鉴权的
	{
		systemRoutes.InitLoginRouter(group, router.Sys.Login) //获取登录信息

	}
	//做鉴权的
	group.Use(middlewares.JWTAuthMiddleware())
	{

		systemRoutes.InitGetUser(group, router.Sys.Login)                       //获取登录信息
		systemRoutes.InitSysProfileRouter(group, router.Sys.Profile)            //个人信息
		systemRoutes.InitSysUserRouter(group, router.Sys.User)                  //用户相关
		systemRoutes.InitSysDeptRouter(group, router.Sys.Dept)                  //部门相关
		systemRoutes.InitSysDictDataRouter(group, router.Sys.DictData)          //数据字典信息
		systemRoutes.InitSysRoleRouter(group, router.Sys.Role)                  //角色相关
		systemRoutes.InitSysPermissionRouter(group, router.Sys.Permission)      //权限相关
		systemRoutes.InitSysDictTypeRouter(group, router.Sys.DictType)          //数据字典属性
		systemRoutes.InitSysPostRouter(group, router.Sys.Post)                  //岗位属性
		monitorRoutes.InitSysUserOnlineRouter(group, router.Monitor.UserOnline) //在线用户监控
		monitorRoutes.InitSysLogininforRouter(group, router.Monitor.Logininfor) //登录用户日志
		monitorRoutes.InitServerRouter(group, router.Monitor.Info)              //服务监控
		//systemRoutes.InitSysConfigRouter(group)      //参数配置
		//genTableRoutes.InitGenTableRouter(group)     //代码生成
		//quartzRoutes.InitJobRouter(group)            //定时任务
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

}

// Cors
// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "*")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
