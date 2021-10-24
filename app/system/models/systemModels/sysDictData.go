package systemModels

import (
	commonModels "baize/app/common/commonModels"
)

type SysDictDataVo struct {
	DictCode  int64   `json:"dictCode,string" db:"dict_code"`
	DictSort  int32   `json:"dictSort" db:"dict_sort"`
	DictLabel string  `json:"dictLabel" db:"dict_label"`
	DictValue string  `json:"dictValue" db:"dict_value"`
	DictType  string  `json:"dictType" db:"dict_type"`
	CssClass  *string `json:"cssClass" db:"css_class"`
	ListClass *string `json:"listClass" db:"list_class"`
	IsDefault string  `json:"isDefault" db:"is_default"`
	Status    string  `json:"status" db:"status"`
	Remark    *string `json:"remark" db:"remark"`
	commonModels.BaseEntity
}

type SysDictDataDQL struct {
	DictType  string `form:"dictType" db:"dict_type"`
	DictLabel string `form:"dictLabel" db:"dict_label"`
	Status    string `form:"status"`
	commonModels.BaseEntityDQL
}

type SysDictDataDML struct {
	DictCode  int64  `json:"dictCode,string" db:"dict_code"`
	DictSort  *int32 `json:"dictSort" db:"dict_sort"`
	DictLabel string `json:"dictLabel" db:"dict_label"`
	DictValue string `json:"dictValue" db:"dict_value"`
	DictType  string `json:"dictType" db:"dict_type"`
	CssClass  string `json:"cssClass" db:"css_class"`
	ListClass string `json:"listClass" db:"list_class"`
	IsDefault string `json:"isDefault" db:"is_default"`
	Status    string `json:"status" db:"status"`
	Remark    string `json:"remark" db:"remark"`
	commonModels.BaseEntityDML
}
