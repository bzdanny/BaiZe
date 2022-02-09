package systemModels

import (
	"baize/app/common/commonModels"
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
	Status    string `form:"status" db:"status"`
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

func SysDictDataListToRows(dictDatas []*SysDictDataVo) (rows [][]string) {
	rows = make([][]string, 0, len(dictDatas)+1)
	row1 := []string{
		"字典类型",
		"字典标签",
		"字典键值",
		"状态",
		"备注",
	}
	rows = append(rows, row1)
	for _, dictData := range dictDatas {
		row := make([]string, 4)
		row[0] = dictData.DictType
		row[1] = dictData.DictLabel
		row[2] = dictData.DictValue
		if dictData.Status == "0" {
			row[3] = "正常"
		} else {
			row[3] = "停用"
		}
		if dictData.Remark != nil {
			row[4] = *dictData.Remark
		}
		rows = append(rows, row)

	}
	return
}
