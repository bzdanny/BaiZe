package taskUtils

import (
	"baize/app/common/redis"
	"baize/app/constant/constants"
	"baize/app/quartz/quartzDao"
	"baize/app/quartz/quartzDao/quartzDaoImpl"
)

var jobDao quartzDao.IJobDao = quartzDaoImpl.GetJobDao()

func GetQuartzCache(invokeTarget string) string {
	s := constants.QuartzKey + invokeTarget
	getString := redis.GetString(s)
	if getString == "" {
		job := jobDao.SelectJobByInvokeTarget(invokeTarget)
		params := job.JobParams
		if params != nil {
			getString = *params
			redis.SetString(s, getString, 0)
		}

	}
	return getString
}
