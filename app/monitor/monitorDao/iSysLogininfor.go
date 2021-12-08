package monitorDao

import (
	"baize/app/monitor/monitorModels"
)

type ILogininforDao interface {
	InserLogininfor(logininfor *monitorModels.Logininfor)
	SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64)
	DeleteLogininforByIds(infoIds []int64)
	CleanLogininfor()
}
