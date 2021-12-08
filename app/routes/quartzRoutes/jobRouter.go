package quartzRoutes

import (
	"baize/app/common/middlewares"
	"baize/app/quartz/controller/jobController"
	"github.com/gin-gonic/gin"
)

func InitJobRouter(router *gin.RouterGroup) {
	job := router.Group("/monitor/job")
	job.GET("/list", middlewares.HasPermission("monitor:job:list"), jobController.JobLIst)
	job.GET("/export", middlewares.HasPermission("monitor:job:export"), jobController.JobExport)
	job.GET("/:jobId", middlewares.HasPermission("monitor:job:query"), jobController.JobGetInfo)
	job.POST("", middlewares.HasPermission("monitor:job:add"), jobController.JobAdd)
	job.PUT("", middlewares.HasPermission("monitor:job:edit"), jobController.JobEdit)
	job.PUT("/changeStatus", middlewares.HasPermission("monitor:job:changeStatus"), jobController.JobChangeStatus)
	job.PUT("/run", middlewares.HasPermission("monitor:job:changeStatus"), jobController.JobRun)
	job.DELETE("/:jobIds", middlewares.HasPermission("monitor:job:remove"), jobController.JobRemove)
}
