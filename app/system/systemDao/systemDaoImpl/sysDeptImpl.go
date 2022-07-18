package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type SysDeptDao struct {
	deptSql string
}

func NewSysDeptDao() *SysDeptDao {
	return &SysDeptDao{
		deptSql: `select d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status, d.del_flag, d.create_by, d.create_time from sys_dept d`,
	}
}

func (sysDeptDao *SysDeptDao) SelectDeptList(db dataUtil.DB, dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo) {
	whereSql := ` where d.del_flag = '0'`
	if dept.ParentId != nil {
		whereSql += " AND parent_id = :parent_id"
	}
	if dept.DeptName != "" {
		whereSql += " AND dept_name like cosncat('%', :dept_name, '%')"
	}
	if dept.Status != "" {
		whereSql += " AND status = :status"
	}
	if dept.DataScope != "" {
		whereSql += " AND " + dept.DataScope
	}
	whereSql += " order by d.parent_id, d.order_num"

	return dataUtil.NamedQueryList(db, list, dept, sysDeptDao.deptSql+whereSql, "", "")

}
func (sysDeptDao *SysDeptDao) SelectDeptById(db dataUtil.DB, deptId int64) (dept *systemModels.SysDeptVo) {
	whereSql := ` where d.dept_id = ?`
	dept = new(systemModels.SysDeptVo)
	err := db.Get(dept, sysDeptDao.deptSql+whereSql, deptId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (sysDeptDao *SysDeptDao) InsertDept(db dataUtil.DB, dept *systemModels.SysDeptAdd) {
	insertSQL := `insert into sys_dept(dept_id,parent_id,dept_name,create_by,create_time,update_by,update_time %s)
					values(:dept_id,:parent_id,:dept_name,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if dept.Ancestors != "" {
		key += ",ancestors"
		value += ",:ancestors"
	}
	if dept.OrderNum != "" {
		key += ",order_num"
		value += ",:order_num"
	}
	if dept.Leader != "" {
		key += ",leader"
		value += ",:leader"
	}
	if dept.Phone != "" {
		key += ",phone"
		value += ",:phone"
	}
	if dept.Email != "" {
		key += ",email"
		value += ",:email"
	}
	if dept.Status != "" {
		key += ",status"
		value += ",:status"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, dept)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDeptDao *SysDeptDao) UpdateDept(db dataUtil.DB, dept *systemModels.SysDeptEdit) {
	updateSQL := `update sys_dept set update_time = now() , update_by = :update_by`

	if dept.ParentId != 0 {
		updateSQL += ",parent_id = :parent_id"
	}

	if dept.DeptName != "" {
		updateSQL += ",dept_name = :dept_name"
	}
	if dept.Ancestors != "" {
		updateSQL += ",ancestors = :ancestors"
	}
	if dept.OrderNum != "" {
		updateSQL += ",order_num = :order_num"
	}
	if dept.Leader != "" {
		updateSQL += ",leader = :leader"
	}
	if dept.Phone != "" {
		updateSQL += ",phone = :phone"
	}
	if dept.Email != "" {
		updateSQL += ",email = :email"
	}
	if dept.Status != "" {
		updateSQL += ",status = :status"
	}

	updateSQL += " where dept_id = :dept_id"

	_, err := db.NamedExec(updateSQL, dept)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDeptDao *SysDeptDao) DeleteDeptById(db dataUtil.DB, deptId int64) {
	_, err := db.Exec("update sys_dept set del_flag = '2' where dept_id =?", deptId)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDeptDao *SysDeptDao) CheckDeptNameUnique(db dataUtil.DB, deptName string, parentId int64) int64 {
	var roleId int64 = 0
	err := db.Get(&roleId, "select dept_id from sys_dept where dept_name=? and parent_id = ?", deptName, parentId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}
func (sysDeptDao *SysDeptDao) HasChildByDeptId(db dataUtil.DB, deptId int64) int {
	var count = 0
	err := db.Get(&count, "select count(1) from sys_dept where parent_id = ?", deptId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
func (sysDeptDao *SysDeptDao) CheckDeptExistUser(db dataUtil.DB, deptId int64) int {
	var count = 0
	err := db.Get(&count, "select count(1) from sys_user where dept_id = ? and del_flag = '0'", deptId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}

func (sysDeptDao *SysDeptDao) SelectDeptListByRoleId(db dataUtil.DB, roleId int64, deptCheckStrictly bool) (deptIds []string) {
	var err error
	deptIds = make([]string, 0)
	sqlstr := `select d.dept_id
		from sys_dept d
            left join sys_role_dept rd on d.dept_id = rd.dept_id
        where rd.role_id = ?`
	if deptCheckStrictly {
		sqlstr += "  and d.dept_id not in (select d.parent_id from sys_dept d inner join sys_role_dept rd on d.dept_id = rd.dept_id and rd.role_id = ?)"
	}
	sqlstr += " order by d.parent_id, d.order_num"
	if deptCheckStrictly {
		err = db.Select(&deptIds, sqlstr, roleId, roleId)
	} else {
		err = db.Select(&deptIds, sqlstr, roleId)
	}

	if err != nil {
		panic(err)
	}
	return
}
