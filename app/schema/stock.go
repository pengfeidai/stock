package schema

type GetStocks struct {
	Code        string `form:"code" json:"code"`
	Name        string `form:"name" json:"name"`
	CurrentPage int    `form:"currentPage" json:"currentPage" binding:"required,gt=0"`
	PageSize    int    `form:"pageSize" json:"pageSize" binding:"required,gt=0"`
}

type UpdateStock struct {
	Code        string `form:"code" json:"code" binding:"required"`
	Performance string `form:"performance" json:"performance`
	Type        string `form:"type" json:"type"`
}
