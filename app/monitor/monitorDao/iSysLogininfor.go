package monitorDao

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type ILogininforDao interface {
	InserLogininfor(db dataUtil.DB, logininfor *monitorModels.Logininfor)
	SelectLogininforList(db dataUtil.DB, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64)
	DeleteLogininforByIds(db dataUtil.DB, infoIds []int64)
	CleanLogininfor(db dataUtil.DB)
}
