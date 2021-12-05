package quartzService

import "baize/app/quartz/quartzModels"

type IJaoService interface {
	SelectJobList(job *quartzModels.JobDQL) (list []*quartzModels.JobVo, total *int64)
	SelectJobById(id int64) (job *quartzModels.JobVo)
	DeleteJob(job *quartzModels.JobVo)
	DeleteJobByIds(jobIds []int64)
	ChangeStatus(job *quartzModels.JobDML)(code int8)
	Run(job *quartzModels.JobVo)
	InsertJob(job *quartzModels.JobDML)
	UpdateJob(job *quartzModels.JobDML)
}
