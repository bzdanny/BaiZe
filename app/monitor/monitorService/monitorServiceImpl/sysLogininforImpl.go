package monitorServiceImpl

import (
	"baize/app/monitor/monitorDao"
	"baize/app/monitor/monitorDao/monitorDaoImpl"
	"baize/app/monitor/monitorModels"
	"baize/app/utils/exceLize"
	"baize/app/utils/snowflake"
)

var logininforImpl = &logininforService{logininforDao: monitorDaoImpl.GetLogininforDao()}

type logininforService struct {
	logininforDao monitorDao.ILogininforDao
}

func GetLogininforService() *logininforService {
	return logininforImpl
}

func (logininforService *logininforService) SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	return logininforService.logininforDao.SelectLogininforList(logininfor)

}
func (logininforService *logininforService) ExportLogininfor(logininfor *monitorModels.LogininforDQL)(data []byte) {
	list, _ := logininforService.logininforDao.SelectLogininforList(logininfor)
	return exceLize.SetRows(monitorModels.SysLogininforToRows(list))

}

func (logininforService *logininforService) InserLogininfor(loginUser *monitorModels.Logininfor) {
	loginUser.InfoId = snowflake.GenID()
	logininforService.logininforDao.InserLogininfor(loginUser)
}

func (logininforService *logininforService) DeleteLogininforByIds(infoIds []int64) {
	logininforService.logininforDao.DeleteLogininforByIds(infoIds)

}

func (logininforService *logininforService) CleanLogininfor() {
	logininforService.logininforDao.CleanLogininfor()

}
