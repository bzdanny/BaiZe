package systemModels

import (
	commonModels "baize/app/common/commonModels"
)

type SysDictTypeVo struct {
	DictId   int64   `json:"dictId,string" db:"dict_id"`
	DictName string  `json:"dictName" db:"dict_name"`
	DictType string  `json:"dictType" db:"dict_type"`
	Status   string  `json:"status" db:"status"`
	Remark   *string `json:"remark" db:"remark"`
	commonModels.BaseEntity
}

type SysDictTypeDQL struct {
	DictName  string `form:"dictName" db:"dict_name"`
	Status    string `form:"status" db:"status"`
	DictType  string `form:"dictType" db:"dict_type"`
	BeginTime string `form:"beginTime" db:"begin_time"`
	EndTime   string `form:"endTime" db:"end_time"`
	commonModels.BaseEntityDQL
}

type SysDictTypeDML struct {
	DictId   int64  `json:"dictId,string" db:"dict_id"`
	DictName string `json:"dictName" db:"dict_name"`
	DictType string `json:"dictType" db:"dict_type"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	commonModels.BaseEntityDML
}
