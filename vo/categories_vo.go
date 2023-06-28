package vo

type CategoryVo struct {
	CategoryName string `json:"category_name"`
	// 父级ID
	ParentId uint `json:"parent_id"`
}
