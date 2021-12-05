package commonModels

type pageDomain struct {
	OrderBy string `form:"order_by"`
	IsAsc   string `form:"is_asc"`
	Page    int64  `form:"pageNum"`
	Size    int64  `form:"pageSize"`
}

func NewPageDomain() *pageDomain {
	return &pageDomain{Page: 1, Size: 10}
}
