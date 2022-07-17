package systemDaoImpl

import (
	"database/sql"
	"fmt"
	"github.com/bzdanny/BaiZe/app/system/models/systemModels"
	"github.com/bzdanny/BaiZe/baize/datasource/dataUtil"
)

type SysMenuDao struct {
	menuSql string
}

func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		menuSql: `select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time from sys_menu m`,
	}
}

func (md *SysMenuDao) SelectMenuById(db dataUtil.DB, menuId int64) (menu *systemModels.SysMenuVo) {
	menu = new(systemModels.SysMenuVo)
	err := db.Get(menu, md.menuSql+` where menu_id = ?`, menuId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (md *SysMenuDao) SelectMenuList(db dataUtil.DB, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo) {
	whereSql := ``
	if menu.MenuName != "" {
		whereSql += " AND menu_name like concat('%', :menu_name, '%')"
	}
	if menu.Visible != "" {
		whereSql += " AND Visible = :visible"
	}
	if menu.MenuName != "" {
		whereSql += "  AND status = :status"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	whereSql += " order by m.parent_id, m.order_num"

	return dataUtil.NamedQueryList(db, list, menu, md.menuSql+whereSql, "", "")

}

func (md *SysMenuDao) SelectMenuListByUserId(db dataUtil.DB, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo) {
	whereSql := ` left join sys_role_menu rm on m.menu_id = rm.menu_id
		left join sys_user_role ur on rm.role_id = ur.role_id
		left join sys_role ro on ur.role_id = ro.role_id
		where ur.user_id = :user_id`
	if menu.MenuName != "" {
		whereSql += " AND menu_name like concat('%', :menu_name, '%')"
	}
	if menu.Visible != "" {
		whereSql += " AND Visible = :visible"
	}
	if menu.MenuName != "" {
		whereSql += "  AND status = :status"
	}
	whereSql += " m.parent_id, m.order_num"
	list = make([]*systemModels.SysMenuVo, 0, 2)

	return dataUtil.NamedQueryList(db, list, menu, md.menuSql+whereSql, "", "")

}

func (md *SysMenuDao) InsertMenu(db dataUtil.DB, menu *systemModels.SysMenuDML) {
	insertSQL := `insert into sys_menu(menu_id,menu_name,parent_id,create_by,create_time,update_by,update_time %s)
					values(:menu_id,:menu_name,:parent_id,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if menu.OrderNum != "" {
		key += ",order_num"
		value += ",:order_num"
	}
	if menu.Path != "" {
		key += ",path"
		value += ",:path"
	}
	if menu.Component != "" {
		key += ",component"
		value += ",:component"
	}
	if menu.IsFrame != "" {
		key += ",is_frame"
		value += ",:is_frame"
	}
	if menu.IsCache != "" {
		key += ",is_cache"
		value += ",:is_cache"
	}
	if menu.MenuType != "" {
		key += ",menu_type"
		value += ",:menu_type"
	}
	if menu.Visible != "" {
		key += ",visible"
		value += ",:visible"
	}
	if menu.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if menu.OrderNum != "" {
		key += ",Perms"
		value += ",:perms"
	}
	if menu.Icon != "" {
		key += ",icon"
		value += ",:icon"
	}
	if menu.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExec(insertStr, menu)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysMenuDao) UpdateMenu(db dataUtil.DB, menu *systemModels.SysMenuDML) {
	updateSQL := `update sys_menu set update_time = now() , update_by = :update_by`

	if menu.ParentId != 0 {
		updateSQL += ",parent_id = :parent_id"
	}
	if menu.MenuName != "" {
		updateSQL += ",menu_name = :menu_name"
	}
	if menu.OrderNum != "" {
		updateSQL += ",order_num = :order_num"
	}
	if menu.Path != "" {
		updateSQL += ",path = :path"
	}
	if menu.Component != "" {
		updateSQL += ",component = :component"
	}
	if menu.IsFrame != "" {
		updateSQL += ",is_frame = :is_frame"
	}
	if menu.IsCache != "" {
		updateSQL += ",is_cache = :is_cache"
	}
	if menu.MenuType != "" {
		updateSQL += ",menu_type = :menu_type"
	}
	if menu.Visible != "" {
		updateSQL += ",visible = :visible"
	}
	if menu.Status != "" {
		updateSQL += ",status = :status"
	}
	if menu.Perms != "" {
		updateSQL += ",perms = :perms"
	}
	if menu.Icon != "" {
		updateSQL += ",icon = :icon"
	}
	if menu.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where menu_id = :menu_id"

	_, err := db.NamedExec(updateSQL, menu)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysMenuDao) DeleteMenuById(db dataUtil.DB, menuId int64) {
	_, err := db.Exec("delete from sys_menu where menu_id = ?", menuId)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysMenuDao) SelectMenuPermsByUserId(db dataUtil.DB, userId int64) (perms []string) {
	sqlStr := `	select distinct m.perms
				from sys_menu m
				left join sys_role_menu rm on m.menu_id = rm.menu_id
				left join sys_user_role ur on rm.role_id = ur.role_id
				left join sys_role r on r.role_id = ur.role_id
				where m.status = '0' and r.status = '0' and ur.user_id =  ?`
	perms = make([]string, 0, 2)
	err := db.Select(&perms, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysMenuDao) SelectMenuTreeAll(db dataUtil.DB) (sysMenus []*systemModels.SysMenuVo) {
	whereSql := ` where m.menu_type in ('M', 'C') and m.status = 0
		order by m.parent_id, m.order_num`
	sysMenus = make([]*systemModels.SysMenuVo, 0, 2)
	err := db.Select(&sysMenus, md.menuSql+whereSql)
	if err != nil {
		panic(err)
	}
	return
}
func (md *SysMenuDao) SelectMenuTreeByUserId(db dataUtil.DB, userId int64) (sysMenus []*systemModels.SysMenuVo) {
	whereSql := ` left join sys_role_menu rm on m.menu_id = rm.menu_id
			 left join sys_user_role ur on rm.role_id = ur.role_id
			 left join sys_role ro on ur.role_id = ro.role_id
			 left join sys_user u on ur.user_id = u.user_id
		where u.user_id = ? and m.menu_type in ('M', 'C') and m.status = 0  AND ro.status = 0
		order by m.parent_id, m.order_num`
	sysMenus = make([]*systemModels.SysMenuVo, 0, 2)
	err := db.Select(&sysMenus, md.menuSql+whereSql, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (md *SysMenuDao) CheckMenuNameUnique(db dataUtil.DB, menuName string, parentId int64) int64 {
	var roleId int64 = 0
	err := db.Get(&roleId, "select menu_id from sys_menu where menu_name=? and parent_id = ?", menuName, parentId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func (md *SysMenuDao) HasChildByMenuId(db dataUtil.DB, menuId int64) int {
	var count = 0
	err := db.Get(&count, "select count(1) from sys_menu where parent_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
func (md *SysMenuDao) SelectMenuListByRoleId(db dataUtil.DB, roleId int64, menuCheckStrictly bool) (roleIds []string) {
	var err error
	roleIds = make([]string, 0, 2)
	sqlstr := `select m.menu_id
		from sys_menu m
            left join sys_role_menu rm on m.menu_id = rm.menu_id
        where rm.role_id = ?`
	if menuCheckStrictly {
		sqlstr += " and m.menu_id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.menu_id = rm.menu_id and rm.role_id = ?)"
		err = db.Select(&roleIds, sqlstr, roleId, roleId)
	} else {
		err = db.Select(&roleIds, sqlstr, roleId)
	}

	if err != nil {
		panic(err)
	}
	return
}
