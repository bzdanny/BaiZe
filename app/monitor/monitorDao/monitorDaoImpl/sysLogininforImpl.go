package monitorDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/constant/constants"
	"baize/app/monitor/monitorModels"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var sysDeptDaoImpl *logininforDao = &logininforDao{db: mysql.GetMysqlDb()}

type logininforDao struct {
	db **sqlx.DB
}

func GetLogininforDao() *logininforDao {
	return sysDeptDaoImpl
}
func (logininforDao *logininforDao) getDb() *sqlx.DB {
	return *logininforDao.db
}

var selectLogininforSql = `select info_id, user_name, ipaddr, login_location, browser, os, status, msg, login_time `
var fromLogininforSql = ` from sys_logininfor`

func (logininforDao *logininforDao) InserLogininfor(logininfor *monitorModels.Logininfor) {

	_, err := logininforDao.getDb().NamedExec("insert into sys_logininfor (info_id,user_name, status, ipaddr, login_location, browser, os, msg, login_time) values (:info_id,:user_name, :status, :ipaddr, :login_location, :browser, :os, :msg, sysdate())", logininfor)
	if err != nil {
		zap.L().Error("登录信息保存错误", zap.Error(err))
	}
	return
}
func (logininforDao *logininforDao) SelectLogininforList(logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	whereSql := ``
	if logininfor.IpAddr != "" {
		whereSql += " AND ipaddr like concat('%', #{ipaddr}, '%')"
	}
	if logininfor.Status != "" {
		whereSql += " AND  status = :status"
	}
	if logininfor.UserName != "" {
		whereSql += " AND user_name like concat('%', #{userName}, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	countRow, err := logininforDao.getDb().NamedQuery(constants.MysqlCount+fromLogininforSql+whereSql, logininfor)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	list = make([]*monitorModels.Logininfor, 0, logininfor.Size)
	if *total > logininfor.Offset {
		if logininfor.Limit != "" {
			whereSql += logininfor.Limit
		}
		listRows, err := logininforDao.getDb().NamedQuery(selectLogininforSql+fromLogininforSql+whereSql, logininfor)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			m := new(monitorModels.Logininfor)
			listRows.StructScan(m)
			list = append(list, m)
		}
		defer listRows.Close()
	}
	return

}
func (logininforDao *logininforDao) DeleteLogininforByIds(infoIds []int64) {
	query, i, err := sqlx.In("delete from sys_logininfor where info_id in (?)", infoIds)
	if err != nil {
		panic(err)
	}
	_, err = logininforDao.getDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}

func (logininforDao *logininforDao) CleanLogininfor() {
	_, err := logininforDao.getDb().Exec("truncate table sys_logininfor")
	if err != nil {
		panic(err)
	}
}
