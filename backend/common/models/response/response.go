package response

type PageResult[T any] struct {
	Entries    []*T  `json:"entries" binding:"required"`                       // 对象列表
	TotalCount int64 `json:"total_count" binding:"required,gte=0" example:"3"` // 当前筛选条件下的对象数量
}

type IDResp struct {
	ID int64 `json:"id" binding:"required" example:"1"` // 对象ID
}

type IDNameResp struct {
	ID   int64  `json:"id" binding:"required" example:"1"`                        // 对象ID
	Name string `json:"name" binding:"required,min=1,max=128" example:"obj_name"` // 对象名称
}
