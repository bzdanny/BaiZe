package monitorModels

import (
	"baize/app/common/commonModels"
)

type JobVo struct {
	JobId          int64   `json:"jobId""`
	JobName        string  `json:"jobName"`
	JobGroup       string  `json:"jobGroup"`
	JobParams      *string `json:"jobParams"`
	InvokeTarget   string  `json:"invokeTarget"`
	CronExpression string  `json:"cronExpression"`
	MisfirePolicy  string  `json:"misfirePolicy"`
	Concurrent     int8    `json:"concurrent"`
	Status         int8    `json:"status"`
	Remark         string  `json:"remark"`
	commonModels.BaseEntity
}
