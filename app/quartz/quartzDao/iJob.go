package quartzDao

import "baize/app/quartz/quartzModels"

type IJobDao interface {
	SelectJobList(job *quartzModels.JobDQL)(list []*quartzModels.JobVo,total *int64)
	SelectJobAll()(list []*quartzModels.JobVo)
	SelectJobById(id int64)(job *quartzModels.JobVo)
	SelectJobByInvokeTarget(invokeTarget string)(job *quartzModels.JobVo)
	DeleteJobById(id int64)
	UpdateJob(job *quartzModels.JobDML)
	InsertJob(job *quartzModels.JobDML)
	DeleteJobByIds(id []int64)
}
