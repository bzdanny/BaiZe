package monitorService

import "baize/app/monitor/monitorModels"

type ISysOperLogService interface {
	InsertOperLog(operLog *monitorModels.SysOpenLog)
	SelectOperLogList(openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64)
	ExportOperLog(logininfor *monitorModels.SysOpenLogDQL) (data []byte)
	DeleteOperLogByIds(operIds []int64)
	SelectOperLogById(operId int64) (operLogList *monitorModels.SysOpenLog)
	CleanOperLog()
}
