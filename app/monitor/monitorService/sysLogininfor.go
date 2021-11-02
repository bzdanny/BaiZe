package monitorService

import (
	"baize/app/monitor/monitorDao"
	"baize/app/monitor/monitorModels"
)

func SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	return monitorDao.SelectLogininforList(logininfor)

}

func InserLogininfor(loginUser *monitorModels.Logininfor) {
	monitorDao.InserLogininfor(loginUser)
}

func DeleteLogininforByIds(infoIds []int64) {
	monitorDao.DeleteLogininforByIds(infoIds)

}

func CleanLogininfor() {
	monitorDao.CleanLogininfor()

}
