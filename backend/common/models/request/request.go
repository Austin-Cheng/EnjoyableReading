package request

type PageInfo struct {
	Offset    int    `json:"offset" form:"offset,default=1" binding:"omitempty,min=1" default:"1"  example:"1"`                                                  // 页码，默认1
	Limit     int    `json:"limit" form:"limit,default=12" binding:"omitempty,min=1,max=120" default:"12"  example:"12"`                                         // 每页大小，默认12
	Direction string `json:"direction" form:"direction,default=desc" binding:"omitempty,oneof=asc desc" default:"desc"  example:"desc"`                          // 排序方向，枚举：asc：正序；desc：倒序。默认倒序
	Sort      string `json:"sort" form:"sort,default=created_at" binding:"omitempty,oneof=created_at updated_at name" default:"created_at" example:"updated_at"` // 排序类型，枚举：created_at：按创建时间排序；updated_at：按更新时间排序; name: 按名称排序。默认按创建时间排序
}

type KeywordInfo struct {
	Keyword string `json:"keyword" form:"keyword" binding:"VerifyXssString,omitempty,min=1,max=128"` // 关键字查询，字符无限制
}

type PageInfoWithKeyword struct {
	PageInfo
	KeywordInfo
}
