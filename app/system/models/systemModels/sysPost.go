package systemModels

import (
	commonModels "baize/app/common/commonModels"
)

type SysPostDQL struct {
	PostCode string `form:"postCode" db:"post_code"`
	Status   string `form:"status" db:"status"`
	PostName string `form:"postName" db:"post_name"`
	commonModels.BaseEntityDQL
}

type SysPostDML struct {
	PostId   int64  `json:"postId,string" db:"post_id"`
	PostCode string `json:"postCode" db:"post_code"`
	PostName string `json:"postName" db:"post_name"`
	PostSort string `json:"postSort" db:"post_sort"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	commonModels.BaseEntityDML
}

type SysPostVo struct {
	PostId   int64   `json:"postId,string" db:"post_id"`
	PostCode string  `json:"postCode" db:"post_code"`
	PostName string  `json:"postName" db:"post_name"`
	PostSort string  `json:"postSort" db:"post_sort"`
	Status   string  `json:"status" db:"status"`
	Remark   *string `json:"remark" db:"remark"`
	commonModels.BaseEntity
}
