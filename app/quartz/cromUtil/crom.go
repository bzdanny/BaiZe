package cromUtil

import (
	"github.com/robfig/cron/v3"
)

var cromMap = make(map[int64]*cron.Cron)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func Start(cmd func(), jobId int64, cronExpression string) {
	c := newWithSeconds()
	_, err := c.AddFunc(cronExpression, cmd)
	if err != nil {
		panic(err)
	}
	c.Start()
	cromMap[jobId] = c

}

func Stop(jobId int64) {
	cromMap[jobId].Stop()
	delete(cromMap, jobId)

}
func IsExistCron(jobId int64)(ok bool){
	_,ok=cromMap[jobId]
	return
}