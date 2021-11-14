package monitorDao

import (
	"baize/app/monitor/monitorModels"
)

type IOperLog interface {
	InsertOperLog(operLog *monitorModels.SysOpenLog)
	SelectOperLogList(openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64)
	DeleteOperLogByIds(operIds []int64)
	SelectOperLogById(operId int64) (operLog *monitorModels.SysOpenLog)
	CleanOperLog()
}
