package monitorDao

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type IOperLog interface {
	InsertOperLog(db dataUtil.DB, operLog *monitorModels.SysOpenLog)
	SelectOperLogList(db dataUtil.DB, openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64)
	DeleteOperLogByIds(db dataUtil.DB, operIds []int64)
	SelectOperLogById(db dataUtil.DB, operId int64) (operLog *monitorModels.SysOpenLog)
	CleanOperLog(db dataUtil.DB)
}
