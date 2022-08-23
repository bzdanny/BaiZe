package monitorDaoImpl

import (
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type LogininforDao struct {
	selectSql string
}

func NewLogininforDao() *LogininforDao {
	return &LogininforDao{
		selectSql: `select info_id, user_name, ipaddr, login_location, browser, os, status, msg, login_time  from sys_logininfor`,
	}
}

func (ld *LogininforDao) InserLogininfor(db dataUtil.DB, logininfor *monitorModels.Logininfor) {

	_, err := db.NamedExec("insert into sys_logininfor (info_id,user_name, status, ipaddr, login_location, browser, os, msg, login_time) values (:info_id,:user_name, :status, :ipaddr, :login_location, :browser, :os, :msg, sysdate())", logininfor)
	if err != nil {
		panic(err)
	}
	return
}
func (ld *LogininforDao) SelectLogininforList(db dataUtil.DB, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	whereSql := ``
	if logininfor.IpAddr != "" {
		whereSql += " AND ipaddr like concat('%', :ipaddr, '%')"
	}
	if logininfor.Status != "" {
		whereSql += " AND  status = :status"
	}
	if logininfor.UserName != "" {
		whereSql += " AND user_name like concat('%', :userName, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	return dataUtil.NamedQueryListAndTotal(db, list, logininfor, ld.selectSql+whereSql, "", "")

}
func (ld *LogininforDao) DeleteLogininforByIds(db dataUtil.DB, infoIds []int64) {
	query, i, err := sqlx.In("delete from sys_logininfor where info_id in (?)", infoIds)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (ld *LogininforDao) CleanLogininfor(db dataUtil.DB) {
	_, err := db.Exec("truncate table sys_logininfor")
	if err != nil {
		panic(err)
	}
}
