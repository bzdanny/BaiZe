package monitorDaoImpl

import (
	"baize/app/common/mysql"
	"baize/app/constant/constants"
	"baize/app/monitor/monitorModels"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var sysOperLogImpl = &operLogDao{
	selectSql: ` select oper_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time`,
	fromSql:   ` from sys_oper_log`,
}

type operLogDao struct {
	selectSql string
	fromSql   string
}

func GetOperLogDao() *operLogDao {
	return sysOperLogImpl
}

func (operLogDao *operLogDao) InsertOperLog(operLog *monitorModels.SysOpenLog) {
	_, err := mysql.GetMasterMysqlDb().NamedExec("insert into sys_oper_log(oper_id,title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time)"+
		"  values (:oper_id,:title, :business_type, :method, :request_method, :operator_type, :oper_name, :dept_name, :oper_url, :oper_ip, :oper_location, :oper_param, :json_result, :status, :error_msg, sysdate())", operLog)
	if err != nil {
		zap.L().Error("登录信息保存错误", zap.Error(err))
	}
	return
}
func (operLogDao *operLogDao) SelectOperLogList(openLog *monitorModels.SysOpenLogDQL) (list []*monitorModels.SysOpenLog, total *int64) {
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

	countRow, err := mysql.GetMasterMysqlDb().NamedQuery(constants.MysqlCount+operLogDao.fromSql+whereSql, openLog)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	list = make([]*monitorModels.SysOpenLog, 0, openLog.Size)
	if *total > openLog.Offset {
		whereSql += " order by oper_id desc"
		if openLog.Limit != "" {
			whereSql += openLog.Limit
		}
		listRows, err := mysql.GetMasterMysqlDb().NamedQuery(operLogDao.selectSql+operLogDao.fromSql+whereSql, openLog)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			m := new(monitorModels.SysOpenLog)
			err := listRows.StructScan(m)
			if err != nil {
				panic(err)
			}
			list = append(list, m)
		}
		defer listRows.Close()
	}
	return
}
func (operLogDao *operLogDao) DeleteOperLogByIds(operIds []int64) {
	query, i, err := sqlx.In("delete from sys_oper_log where oper_id in (?)", operIds)
	if err != nil {
		panic(err)
	}
	_, err = mysql.GetMasterMysqlDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
func (operLogDao *operLogDao) SelectOperLogById(operId int64) (operLog *monitorModels.SysOpenLog) {
	whereSql := `  where oper_id = ?`
	operLog = new(monitorModels.SysOpenLog)
	err := mysql.GetMasterMysqlDb().Get(operLog, operLogDao.selectSql+operLogDao.fromSql+whereSql, operId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *operLogDao) CleanOperLog() {
	_, err := mysql.GetMasterMysqlDb().Exec("truncate table sys_oper_log")
	if err != nil {
		panic(err)
	}
}
