package base

type Response struct {
	Total    int         `json:"total"`    // 总条数
	Page     int         `json:"page"`     // 页数
	PageSize int         `json:"pageSize"` // 每页条数
	Data     interface{} `json:"data"`     // 数据
}
