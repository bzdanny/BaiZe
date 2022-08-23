package systemModels

import "github.com/bzdanny/BaiZe/baize/baizeEntity"

type SysDictTypeVo struct {
	DictId   int64   `json:"dictId,string" db:"dict_id"`
	DictName string  `json:"dictName" db:"dict_name"`
	DictType string  `json:"dictType" db:"dict_type"`
	Status   string  `json:"status" db:"status"`
	Remark   *string `json:"remark" db:"remark"`
	baizeEntity.BaseEntity
}

type SysDictTypeDQL struct {
	DictName  string `form:"dictName" db:"dict_name"`
	Status    string `form:"status" db:"status"`
	DictType  string `form:"dictType" db:"dict_type"`
	BeginTime string `form:"beginTime" db:"begin_time"`
	EndTime   string `form:"endTime" db:"end_time"`
	baizeEntity.BaseEntityDQL
}

type SysDictTypeAdd struct {
	DictId   int64  `json:"dictId,string" db:"dict_id"`
	DictName string `json:"dictName" db:"dict_name"`
	DictType string `json:"dictType" db:"dict_type"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	baizeEntity.BaseEntityAdd
}
type SysDictTypeEdit struct {
	DictId   int64  `json:"dictId,string" db:"dict_id"`
	DictName string `json:"dictName" db:"dict_name"`
	DictType string `json:"dictType" db:"dict_type"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	baizeEntity.BaseEntityEdit
}

func SysDictTypeListToRows(dictTypes []*SysDictTypeVo) (rows [][]string) {
	rows = make([][]string, 0, len(dictTypes)+1)
	row1 := []string{"字典名称", "字典类型", "状态", "备注"}
	rows = append(rows, row1)
	for _, dictType := range dictTypes {
		row := make([]string, 4)
		row[0] = dictType.DictName
		row[1] = dictType.DictType
		if dictType.Status == "0" {
			row[2] = "正常"
		} else {
			row[2] = "停用"
		}
		if dictType.Remark != nil {
			row[3] = *dictType.Remark
		}
		rows = append(rows, row)

	}
	return
}
