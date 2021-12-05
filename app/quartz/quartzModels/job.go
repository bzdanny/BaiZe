package quartzModels

import (
	"baize/app/common/commonModels"
)

type JobVo struct {
	JobId          int64   `json:"jobId,string" db:"job_id"`
	JobName        string  `json:"jobName" db:"job_name"`
	JobGroup       string  `json:"jobGroup" db:"job_group"`
	JobParams      *string `json:"jobParams" db:"job_params"`
	InvokeTarget   string  `json:"invokeTarget" db:"invoke_target"`
	CronExpression string  `json:"cronExpression" db:"cron_expression"`
	Status         string    `json:"status" db:"status"`
	Remark         string  `json:"remark" db:"remark"`
	commonModels.BaseEntity
}
type JobDQL struct {
	JobName  string `form:"jobName" db:"job_name"`
	JobGroup string `form:"jobGroup" db:"job_group"`
	InvokeTarget string `form:"invokeTarget" db:"invoke_target"`
	Status   *string   `form:"status" db:"Status"`
	commonModels.BaseEntityDQL
}

type JobDML struct {
	JobId          int64   `json:"jobId,string" db:"job_id"`
	JobName        *string `json:"jobName" db:"job_name"`
	JobGroup       *string `json:"jobGroup" db:"job_group"`
	JobParams      *string `json:"jobParams" db:"job_params"`
	InvokeTarget   *string `json:"invokeTarget" db:"invoke_target"`
	Concurrent     *int8   `json:"concurrent" db:"concurrent"`
	Status         *string   `json:"status" db:"status"`
	Remark         *string `json:"remark" db:"remark"`
	commonModels.BaseEntityDML
}
