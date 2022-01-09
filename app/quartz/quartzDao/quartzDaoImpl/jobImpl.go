package quartzDaoImpl

import (
	"baize/app/common/datasource"
	"baize/app/constant/constants"
	"baize/app/quartz/quartzModels"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var jobDaoDaoImpl *jobDao

func init() {
	jobDaoDaoImpl = &jobDao{
		selectSql: "select job_id, job_name,job_params, job_group, invoke_target, cron_expression,status, create_by, create_time, remark ",
		fromSql:   " from sys_job",
	}
}

type jobDao struct {
	selectSql string
	fromSql   string
}

func GetJobDao() *jobDao {
	return jobDaoDaoImpl
}

func (jobDao *jobDao) SelectJobList(job *quartzModels.JobDQL) (list []*quartzModels.JobVo, total *int64) {
	whereSql := ``
	if job.JobName != "" {
		whereSql += " AND job_name like concat('%', :job_name, '%')"
	}
	if job.JobGroup != "" {
		whereSql += " AND job_group = :job_group"
	}
	if job.Status != nil {
		whereSql += " AND status = :status"
	}
	if job.InvokeTarget != "" {
		whereSql += " AND invoke_target like concat('%', :invoke_target, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	countRow, err := datasource.GetMasterDb().NamedQuery(constants.MysqlCount+jobDao.fromSql+whereSql, job)
	if err != nil {
		panic(err)
	}
	total = new(int64)
	if countRow.Next() {
		countRow.Scan(total)
	}
	defer countRow.Close()
	list = make([]*quartzModels.JobVo, 0, job.Size)
	if *total > job.Offset {
		if job.Limit != "" {
			whereSql += job.Limit
		}
		listRows, err := datasource.GetMasterDb().NamedQuery(jobDao.selectSql+jobDao.fromSql+whereSql, job)
		if err != nil {
			panic(err)
		}
		for listRows.Next() {
			m := new(quartzModels.JobVo)
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
func (jobDao *jobDao) SelectJobAll() (list []*quartzModels.JobVo) {
	list = make([]*quartzModels.JobVo, 0)
	err := datasource.GetMasterDb().Select(&list, jobDao.selectSql+jobDao.fromSql)
	if err != nil {
		panic(err)
	}
	return
}
func (jobDao *jobDao) SelectJobById(id int64) (job *quartzModels.JobVo) {
	job = new(quartzModels.JobVo)
	err := datasource.GetMasterDb().Get(job, jobDao.selectSql+jobDao.fromSql+" where job_id = ?", id)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (jobDao *jobDao) SelectJobByInvokeTarget(invokeTarget string) (job *quartzModels.JobVo) {
	job = new(quartzModels.JobVo)
	err := datasource.GetMasterDb().Get(job, jobDao.selectSql+jobDao.fromSql+" where invoke_target = ?", invokeTarget)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (jobDao *jobDao) DeleteJobById(id int64) {
	_, err := datasource.GetMasterDb().Exec("delete from sys_job where job_id =", id)
	if err != nil {
		panic(err)
	}
	return
}
func (jobDao *jobDao) UpdateJob(job *quartzModels.JobDML) {
	updateSQL := `update sys_job set update_time = now() , update_by = :update_by`

	if job.JobName != nil {
		updateSQL += ",job_name = :job_name"
	}

	if job.JobGroup != nil {
		updateSQL += ",job_group = :job_group"
	}

	if job.JobParams != nil {
		updateSQL += ",job_params = :job_params"
	}

	if job.InvokeTarget != nil {
		updateSQL += ",invoke_target = :invoke_target"
	}

	if job.Concurrent != nil {
		updateSQL += ",concurrent = :concurrent"
	}

	if job.Status != nil {
		updateSQL += ",status = :status"
	}
	if job.Remark != nil {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where job_id = :job_id"

	_, err := datasource.GetMasterDb().NamedExec(updateSQL, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jobDao *jobDao) InsertJob(job *quartzModels.JobDML) {
	insertSQL := `insert into sys_job(job_id,job_name,job_group,invoke_target,cron_expression,,status,create_by,create_time,update_by,update_time %s)
					values(:job_id,:job_name,:job_group,:invoke_target,:cron_expression,:status,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if job.JobParams != nil {
		key += ",job_group"
		value += ",:job_group"
	}
	if job.Remark != nil {
		key += ",remark"
		value += ",:remark"
	}
	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := datasource.GetMasterDb().NamedExec(insertStr, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jobDao *jobDao) DeleteJobByIds(ids []int64) {
	query, i, err := sqlx.In("delete from sys_job where job_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = datasource.GetMasterDb().Exec(query, i...)
	if err != nil {
		panic(err)
	}
}
