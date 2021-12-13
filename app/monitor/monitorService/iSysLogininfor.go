package monitorService

import (
	"baize/app/monitor/monitorModels"
)

type ILogininforService interface {
	SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64)
	ExportLogininfor(logininfor *monitorModels.LogininforDQL)(data []byte)
	InserLogininfor(loginUser *monitorModels.Logininfor)
	DeleteLogininforByIds(infoIds []int64)
	CleanLogininfor()
}
