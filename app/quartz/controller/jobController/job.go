package jobController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/quartz/quartzModels"
	"baize/app/quartz/quartzService"
	"baize/app/quartz/quartzService/quartzServiceImpl"
	"github.com/gin-gonic/gin"
)

var iJob quartzService.IJaoService = quartzServiceImpl.GetLogininforService()

func JobLIst(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	job := new(quartzModels.JobDQL)
	c.ShouldBind(job)
	job.SetLimit(c)
	list, total := iJob.SelectJobList(job)
	bzc.SuccessListData(list, total)
}
func JobExport(c *gin.Context) {

}
func JobGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	jobId := bzc.ParamInt64("jobId")
	menu := iJob.SelectJobById(jobId)
	bzc.SuccessData(menu)
}
func JobAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("定时任务", "INSERT")
	job := new(quartzModels.JobDML)
	c.ShouldBind(job)
	job.SetCreateBy(bzc.GetCurrentUserName())
	iJob.InsertJob(job)
	bzc.Success()
}
func JobEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("定时任务", "UPDATE")
	job := new(quartzModels.JobDML)
	c.ShouldBind(job)
	job.SetCreateBy(bzc.GetCurrentUserName())
	iJob.UpdateJob(job)
	bzc.Success()
}
func JobChangeStatus(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("定时任务", "UPDATE")
	job := new(quartzModels.JobDML)
	c.ShouldBind(job)
	status := iJob.ChangeStatus(job)
	if status == 1 {
		bzc.Waring("目标方法未找到")
		return
	}
	bzc.Success()
}
func JobRun(c *gin.Context) {

}
func JobRemove(c *gin.Context) {

}
