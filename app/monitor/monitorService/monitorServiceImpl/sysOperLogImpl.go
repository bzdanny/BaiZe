package monitorServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorDao"
	"github.com/bzdanny/BaiZe/app/monitor/monitorDao/monitorDaoImpl"
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
)

type OperLogService struct {
	data *datasource.Data
	od   monitorDao.IOperLog
}

func NewOperLogServiceService(data *datasource.Data, od *monitorDaoImpl.OperLogDao) *OperLogService {
	return &OperLogService{
		data: data,
		od:   od,
	}
}

func (ols *OperLogService) InsertOperLog(operLog *monitorModels.SysOpenLog) {
	operLog.OperId = snowflake.GenID()
	ols.od.InsertOperLog(ols.data.GetMasterDb(), operLog)
}
func (ols *OperLogService) SelectOperLogList(openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64) {
	list, total = ols.od.SelectOperLogList(ols.data.GetSlaveDb(), openLog)
	return

}
func (ols *OperLogService) ExportOperLog(openLog *monitorModels.SysOpenLogDQL) (data []byte) {
	//list, _ := ols.od.SelectOperLogList(ols.data.GetSlaveDb(), openLog)
	//return exceLize.SetRows(monitorModels.SysOperLogListToRows(list))
	return nil
}

func (ols *OperLogService) DeleteOperLogByIds(operIds []int64) {
	ols.od.DeleteOperLogByIds(ols.data.GetMasterDb(), operIds)
}
func (ols *OperLogService) SelectOperLogById(operId int64) (operLog *monitorModels.SysOpenLog) {
	operLog = ols.od.SelectOperLogById(ols.data.GetSlaveDb(), operId)
	return
}
func (ols *OperLogService) CleanOperLog() {
	ols.od.CleanOperLog(ols.data.GetMasterDb())
}
