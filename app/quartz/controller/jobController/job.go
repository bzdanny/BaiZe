package jobController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/quartz/quartzModels"
	"baize/app/quartz/quartzService"
	"baize/app/quartz/quartzService/quartzServiceImpl"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var iJob quartzService.IJaoService=quartzServiceImpl.GetLogininforService()

func JobLIst(c *gin.Context)  {
	job := new(quartzModels.JobDQL)
	c.ShouldBind(job)
	job.SetLimit(c)
	list, total := iJob.SelectJobList(job)
	c.JSON(http.StatusOK,commonModels.SuccessListData(list,total))
}
func JobExport(c *gin.Context)  {

}
func JobGetInfo(c *gin.Context)  {
	jobId, err := strconv.ParseInt(c.Param("jobId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
	}
	menu := iJob.SelectJobById(jobId)
	c.JSON(http.StatusOK, commonModels.SuccessData(menu))
}
func JobAdd(c *gin.Context)  {
	commonLog.SetLog(c, "定时任务", "INSERT")
	loginUser := commonController.GetCurrentLoginUser(c)
	job := new(quartzModels.JobDML)
	c.ShouldBind(job)
	job.SetCreateBy(loginUser.User.UserName)
	iJob.InsertJob(job)
	c.JSON(http.StatusOK, commonModels.Success())
}
func JobEdit(c *gin.Context)  {
	commonLog.SetLog(c, "定时任务", "UPDATE")
	loginUser := commonController.GetCurrentLoginUser(c)
	job := new(quartzModels.JobDML)
	c.ShouldBind(job)
	job.SetCreateBy(loginUser.User.UserName)
	iJob.UpdateJob(job)
	c.JSON(http.StatusOK, commonModels.Success())
}
func JobChangeStatus(c *gin.Context)  {
	commonLog.SetLog(c, "定时任务", "UPDATE")
	job := new(quartzModels.JobDML)
	c.ShouldBind(job)
	status := iJob.ChangeStatus(job)
	if status==1 {
		c.JSON(http.StatusOK, commonModels.Waring("目标方法未找到"))
		return
	}
	c.JSON(http.StatusOK, commonModels.Success())
}
func JobRun(c *gin.Context)  {

}
func JobRemove(c *gin.Context)  {

}
