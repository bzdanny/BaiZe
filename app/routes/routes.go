package routes

import (
	"github.com/bzdanny/BaiZe/app/routes/systemRouter"
	"github.com/bzdanny/BaiZe/app/setting"
	"github.com/bzdanny/BaiZe/app/system/systemController"
	"github.com/bzdanny/BaiZe/baize/middlewares"
	"github.com/bzdanny/BaiZe/baize/utils/logger"
	"github.com/google/wire"

	"fmt"
	"net/http"
	"strings"

	_ "github.com/bzdanny/BaiZe/docs"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

var ProviderSet = wire.NewSet(NewRouter)

type Router struct {
	Sys *systemController.SystemController
}

func NewRouter(sys *systemController.SystemController) *Router {
	return &Router{
		Sys: sys,
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
	group.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//不做鉴权的
	{
		systemRouter.InitLoginRouter(group, router.Sys.Login) //获取登录信息

	}
	//做鉴权的
	group.Use(middlewares.JWTAuthMiddleware())
	{

		systemRouter.InitGetUser(group, router.Sys.Login)              //获取登录信息
		systemRouter.InitSysProfileRouter(group, router.Sys.Profile)   //个人信息
		systemRouter.InitSysUserRouter(group, router.Sys.User)         //用户相关
		systemRouter.InitSysDeptRouter(group, router.Sys.Dept)         //部门相关
		systemRouter.InitSysDictDataRouter(group, router.Sys.DictData) //数据字典信息
		systemRouter.InitSysRoleRouter(group, router.Sys.Role)         //角色相关
		systemRouter.InitSysMenuRouter(group, router.Sys.Menu)         //菜单相关
		//systemRouter.InitSysConfigRouter(group)      //参数配置
		systemRouter.InitSysDictTypeRouter(group, router.Sys.DictType) //数据字典属性
		systemRouter.InitSysPostRouter(group, router.Sys.Post)         //岗位属性
		//monitorRoutes.InitSysUserOnlineRouter(group) //在线用户监控
		//monitorRoutes.InitSysLogininforRouter(group) //登录用户日志
		//monitorRoutes.InitSysOperLogRouter(group)    //操作日志
		//monitorRoutes.InitServerRouter(group)        //服务监控
		//genTableRoutes.InitGenTableRouter(group)     //代码生成
		//quartzRoutes.InitJobRouter(group)            //定时任务
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
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
