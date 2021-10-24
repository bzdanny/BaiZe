package systemService

import (
	"baize/app/system/dao/systemDao"
	"baize/app/system/models/systemModels"
)

func SelectLogininforList(logininfor *systemModels.LogininforDQL) (list []*systemModels.Logininfor, total *int64) {
	return systemDao.SelectLogininforList(logininfor)

}

func InserLogininfor(loginUser *systemModels.Logininfor) {
	systemDao.InserLogininfor(loginUser)
}

func DeleteLogininforByIds(infoIds []int64) {
	systemDao.DeleteLogininforByIds(infoIds)

}

func CleanLogininfor() {
	systemDao.CleanLogininfor()

}
