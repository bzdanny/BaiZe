package monitorServiceImpl

import (
	"baize/app/monitor/monitorDao"
	"baize/app/monitor/monitorDao/monitorDaoImpl"
	"baize/app/monitor/monitorModels"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
)

var operLogImpl = &operLogService{operLogDao: monitorDaoImpl.GetOperLogDao()}

type operLogService struct {
	operLogDao monitorDao.IOperLog
}

func GetOperLogServiceService() *operLogService {
	return operLogImpl
}

func (operLogService *operLogService) InsertOperLog(operLog *monitorModels.SysOpenLog) {
	operLog.OperId = snowflake.GenID()
	go operLogService.operLogDao.InsertOperLog(operLog)
}
func (operLogService *operLogService) SelectOperLogList(openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64) {
	list, total = operLogService.operLogDao.SelectOperLogList(openLog)
	return

}
func (operLogService *operLogService) ExportOperLog(openLog *monitorModels.SysOpenLogDQL) (data []byte) {
	list, _ := operLogService.operLogDao.SelectOperLogList(openLog)
	return exceLize.SetRows(monitorModels.SysOperLogListToRows(list))
}

func (operLogService *operLogService) DeleteOperLogByIds(operIds []int64) {
	operLogService.operLogDao.DeleteOperLogByIds(operIds)
}
func (operLogService *operLogService) SelectOperLogById(operId int64) (operLog *monitorModels.SysOpenLog) {
	operLog = operLogService.operLogDao.SelectOperLogById(operId)
	return
}
func (operLogService *operLogService) CleanOperLog() {
	operLogService.operLogDao.CleanOperLog()
}
