package monitorServiceImpl

import (
	"baize/app/monitor/monitorDao"
	"baize/app/monitor/monitorDao/monitorDaoImpl"
	"baize/app/monitor/monitorModels"
)

var logininforImpl *logininforService = &logininforService{logininforDao: monitorDaoImpl.GetLogininforDao()}

type logininforService struct {
	logininforDao monitorDao.ILogininforDao
}

func GetLogininforService() *logininforService {
	return logininforImpl
}

func (logininforService *logininforService) SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	return logininforService.logininforDao.SelectLogininforList(logininfor)

}

func (logininforService *logininforService) InserLogininfor(loginUser *monitorModels.Logininfor) {
	logininforService.logininforDao.InserLogininfor(loginUser)
}

func (logininforService *logininforService) DeleteLogininforByIds(infoIds []int64) {
	logininforService.logininforDao.DeleteLogininforByIds(infoIds)

}

func (logininforService *logininforService) CleanLogininfor() {
	logininforService.logininforDao.CleanLogininfor()

}
