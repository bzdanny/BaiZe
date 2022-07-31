package monitorServiceImpl

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorDao"
	"github.com/bzdanny/BaiZe/app/monitor/monitorDao/monitorDaoImpl"
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/datasource"
	"github.com/bzdanny/BaiZe/pkg/snowflake"
)

type LogininforService struct {
	data *datasource.Data
	ld   monitorDao.ILogininforDao
}

func GetLogininforService(data *datasource.Data, ld *monitorDaoImpl.LogininforDao) *LogininforService {
	return &LogininforService{data: data, ld: ld}
}

func (ls *LogininforService) SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	return ls.ld.SelectLogininforList(ls.data.GetSlaveDb(), logininfor)

}
func (ls *LogininforService) ExportLogininfor(logininfor *monitorModels.LogininforDQL) (data []byte) {
	//list, _ := ls.ld.SelectLogininforList(ls.data.GetSlaveDb(), logininfor)
	//return exceLize.SetRows(monitorModels.SysLogininforToRows(list))
	return nil
}

func (ls *LogininforService) InserLogininfor(loginUser *monitorModels.Logininfor) {
	loginUser.InfoId = snowflake.GenID()
	ls.ld.InserLogininfor(ls.data.GetMasterDb(), loginUser)
}

func (ls *LogininforService) DeleteLogininforByIds(infoIds []int64) {
	ls.ld.DeleteLogininforByIds(ls.data.GetMasterDb(), infoIds)

}

func (ls *LogininforService) CleanLogininfor() {
	ls.ld.CleanLogininfor(ls.data.GetMasterDb())

}
