package monitorDaoImpl

import (
	"database/sql"
	"github.com/bzdanny/BaiZe/app/monitor/monitorModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
	"github.com/jmoiron/sqlx"
)

type OperLogDao struct {
	selectSql string
}

func NewOperLogDao() *OperLogDao {
	return &OperLogDao{
		selectSql: ` select oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time from sys_oper_log`,
	}
}

func (operLogDao *OperLogDao) InsertOperLog(db dataUtil.DB, operLog *monitorModels.SysOpenLog) {
	_, err := db.NamedExec("insert into sys_oper_log(oper_id,title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time)"+
		"  values (:oper_id,:title, :business_type, :method, :request_method, :operator_type, :oper_name, :dept_name, :oper_url, :oper_ip, :oper_location, :oper_param, :json_result, :status, :error_msg, sysdate())", operLog)
	if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *OperLogDao) SelectOperLogList(db dataUtil.DB, openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64) {
	whereSql := ``
	if openLog.Title != "" {
		whereSql += " AND title like concat('%', :title, '%')"
	}
	if openLog.BusinessType != nil {
		whereSql += " AND business_type = :business_type"
	}
	if openLog.Status != nil {
		whereSql += " AND status = :status"
	}
	if openLog.OperName != "" {
		whereSql += " AND oper_name like concat('%', :oper_name, '%')"
	}
	if openLog.BeginTime != "" {
		whereSql += " and date_format(oper_time,'%y%m%d') >= :begin_time"
	}
	if openLog.EndTime != "" {
		whereSql += " and date_format(oper_time,'%y%m%d') >= :end_time"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	return dataUtil.NamedQueryListAndTotal(db, list, openLog, operLogDao.selectSql+whereSql, "", "")
}
func (operLogDao *OperLogDao) DeleteOperLogByIds(db dataUtil.DB, operIds []int64) {
	query, i, err := sqlx.In("delete from sys_oper_log where oper_id in (?)", operIds)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
func (operLogDao *OperLogDao) SelectOperLogById(db dataUtil.DB, operId int64) (operLog *monitorModels.SysOpenLog) {
	operLog = new(monitorModels.SysOpenLog)
	err := db.Get(operLog, operLogDao.selectSql+`  where oper_id = ?`, operId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *OperLogDao) CleanOperLog(db dataUtil.DB) {
	_, err := db.Exec("truncate table sys_oper_log")
	if err != nil {
		panic(err)
	}
}
