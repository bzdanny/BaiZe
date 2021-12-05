package quartzServiceImpl

import (
	"baize/app/quartz/cromUtil"
	"baize/app/quartz/quartzConstant"
	"baize/app/quartz/quartzDao"
	"baize/app/quartz/quartzDao/quartzDaoImpl"
	"baize/app/quartz/quartzModels"
	"baize/app/quartz/task"
	"baize/app/utils/snowflake"
)

var jobImpl = &jobService{jobDao: quartzDaoImpl.GetJobDao()}

type jobService struct {
	jobDao quartzDao.IJobDao
}

func GetLogininforService() *jobService {
	return jobImpl
}

func (jobService *jobService) SelectJobList(job *quartzModels.JobDQL) (list []*quartzModels.JobVo, total *int64) {
	list, total = jobService.jobDao.SelectJobList(job)
	return
}
func (jobService *jobService) SelectJobById(id int64) (job *quartzModels.JobVo) {
	job = jobService.jobDao.SelectJobById(id)
	return
}


func (jobService *jobService) DeleteJob(job *quartzModels.JobVo) {}
func (jobService *jobService) DeleteJobByIds(jobIds []int64)     {}
func (jobService *jobService) ChangeStatus(job *quartzModels.JobDML) (code int8){
	jobVo := jobService.SelectJobById(job.JobId)
	if *job.Status==quartzConstant.Normal {
		if !task.IsExistFunc(jobVo.InvokeTarget){
			code=1
			return
		}
		f := task.GetByName(jobVo.InvokeTarget)
		cromUtil.Start(f,job.JobId,jobVo.CronExpression)

	}else if *job.Status==quartzConstant.Pause {
		if cromUtil.IsExistCron(job.JobId){
			cromUtil.Stop(job.JobId)
		}
	}
	jobService.jobDao.UpdateJob(job)
	return
}
func (jobService *jobService) Run(job *quartzModels.JobVo)        {
	jobVo := jobService.SelectJobById(job.JobId)
	f := task.GetByName(jobVo.InvokeTarget)
	go f()
}
func (jobService *jobService) InsertJob(job *quartzModels.JobDML) {
	job.JobId = snowflake.GenID()
	jobService.jobDao.InsertJob(job)
}
func (jobService *jobService) UpdateJob(job *quartzModels.JobDML) {
	jobService.jobDao.UpdateJob(job)
}
