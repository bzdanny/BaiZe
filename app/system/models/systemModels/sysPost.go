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
	PostSort *int64 `json:"postSort" db:"post_sort"`
	PostCode string `json:"postCode" db:"post_code"`
	PostName string `json:"postName" db:"post_name"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	commonModels.BaseEntityDML
}

type SysPostVo struct {
	PostId   int64   `json:"postId,string" db:"post_id"`
	PostSort *int64  `json:"postSort" db:"post_sort"`
	PostCode string  `json:"postCode" db:"post_code"`
	PostName string  `json:"postName" db:"post_name"`
	Status   string  `json:"status" db:"status"`
	Remark   *string `json:"remark" db:"remark"`
	commonModels.BaseEntity
}
func SysPostListToRows(posts []*SysPostVo) (rows [][]string) {
	rows = make([][]string, 0, len(posts)+1)
	row1 := []string{"岗位编码", "岗位名称", "状态"}
	rows = append(rows, row1)
	for _, post := range posts {
		row := make([]string, 7)
		row[0]=post.PostCode
		row[1]=post.PostName
		if post.Status == "0" {
			row[3] = "正常"
		} else {
			row[3] = "停用"
		}
		rows = append(rows, row)

	}
	return
}