package systemModels

import (
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
)

type SysPermissionDML struct {
	PermissionId   int64  `json:"permissionId,string" db:"permission_id"`
	ParentId       int64  `json:"parentId,string" db:"parent_id"`
	PermissionName string `json:"permissionName" db:"permission_name"`
	Status         string `json:"status" db:"status"`
	Perms          string `json:"perms" db:"perms"`
	Remark         string `json:"remark" db:"remark"`
	baizeEntity.BaseEntityAdd
}

type SysPermissionDQL struct {
	UserId         int64  `db:"userId"`
	PermissionName string `form:"permissionName" db:"permission_name"`
	Status         string `form:"status" db:"status"`
	baizeEntity.BaseEntityDQL
}

type SysPermissionVo struct {
	PermissionId   int64   `json:"permissionId,string" db:"permission_id"`
	PermissionName string  `json:"permissionName" db:"permission_name"`
	ParentName     string  `json:"parentName" db:"parent_name"`
	ParentId       int64   `json:"parentId" db:"parent_id"`
	Status         string  `json:"status" db:"status"`
	Perms          string  `json:"perms" db:"perms"`
	Remark         *string `json:"remark" db:"remark"`
	Children       []*SysPermissionVo
	baizeEntity.BaseEntity
}
