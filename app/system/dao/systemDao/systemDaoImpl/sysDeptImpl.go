package systemDaoImpl

import (
	"baize/app/common/datasource"
	"baize/app/system/models/systemModels"
	"database/sql"
	"fmt"
)

var sysDeptDaoImpl *sysDeptDao

type sysDeptDao struct {
	deptSql string
}

func init() {
	sysDeptDaoImpl = &sysDeptDao{
		deptSql: ` select d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status, d.del_flag, d.create_by, d.create_time from sys_dept d `,
	}
}

func GetSysDeptDao() *sysDeptDao {
	return sysDeptDaoImpl
}

func (deptDao *sysDeptDao) SelectDeptList(dept *systemModels.SysDeptDQL) (sysDeptList []*systemModels.SysDeptVo) {
	fromSql := ` where d.del_flag = '0'`
	if dept.ParentId != nil {
		fromSql += " AND parent_id = :parent_id"
	}
	if dept.DeptName != "" {
		fromSql += " AND dept_name like cosncat('%', :dept_name, '%')"
	}
	if dept.Status != "" {
		fromSql += " AND status = :status"
	}
	if dept.DataScope != "" {
		fromSql += " AND " + dept.DataScope
	}
	fromSql += " order by d.parent_id, d.order_num"
	sysDeptList = make([]*systemModels.SysDeptVo, 0, 2)
	listRows, err := datasource.GetMasterDb().NamedQuery(deptDao.deptSql+fromSql, dept)
	if err != nil {
		panic(err)
	}
	for listRows.Next() {
		sysDept := new(systemModels.SysDeptVo)
		err := listRows.StructScan(sysDept)
		if err != nil {
			panic(err)
		}
		sysDeptList = append(sysDeptList, sysDept)
	}
	return
}
func (deptDao *sysDeptDao) SelectDeptById(deptId int64) (dept *systemModels.SysDeptVo) {
	whereSql := ` where d.dept_id = ?`
	dept = new(systemModels.SysDeptVo)
	err := datasource.GetMasterDb().Get(dept, deptDao.deptSql+whereSql, deptId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (deptDao *sysDeptDao) InsertDept(dept *systemModels.SysDeptDML) {
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
	_, err := datasource.GetMasterDb().NamedExec(insertStr, dept)
	if err != nil {
		panic(err)
	}
	return
}

func (deptDao *sysDeptDao) UpdateDept(dept *systemModels.SysDeptDML) {
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

	_, err := datasource.GetMasterDb().NamedExec(updateSQL, dept)
	if err != nil {
		panic(err)
	}
	return
}

func (deptDao *sysDeptDao) DeleteDeptById(deptId int64) {
	_, err := datasource.GetMasterDb().Exec("update sys_dept set del_flag = '2' where dept_id =?", deptId)
	if err != nil {
		panic(err)
	}
	return
}
func (deptDao *sysDeptDao) CheckDeptNameUnique(deptName string, parentId int64) int64 {
	var roleId int64 = 0
	err := datasource.GetMasterDb().Get(&roleId, "select dept_id from sys_dept where dept_name=? and parent_id = ?", deptName, parentId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}
func (deptDao *sysDeptDao) HasChildByDeptId(deptId int64) int {
	var count = 0
	err := datasource.GetMasterDb().Get(&count, "select count(1) from sys_dept where parent_id = ?", deptId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
func (deptDao *sysDeptDao) CheckDeptExistUser(deptId int64) int {
	var count = 0
	err := datasource.GetMasterDb().Get(&count, "select count(1) from sys_user where dept_id = ? and del_flag = '0'", deptId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}

func (deptDao *sysDeptDao) SelectDeptListByRoleId(roleId int64, deptCheckStrictly bool) (deptIds []string) {
	var err error
	deptIds = make([]string, 0, 2)
	sqlstr := `select d.dept_id
		from sys_dept d
            left join sys_role_dept rd on d.dept_id = rd.dept_id
        where rd.role_id = ?`
	if deptCheckStrictly {
		sqlstr += "  and d.dept_id not in (select d.parent_id from sys_dept d inner join sys_role_dept rd on d.dept_id = rd.dept_id and rd.role_id = ?)"
	}
	sqlstr += " order by d.parent_id, d.order_num"
	if deptCheckStrictly {
		err = datasource.GetMasterDb().Select(&deptIds, sqlstr, roleId, roleId)
	} else {
		err = datasource.GetMasterDb().Select(&deptIds, sqlstr, roleId)
	}

	if err != nil {
		panic(err)
	}
	return
}
