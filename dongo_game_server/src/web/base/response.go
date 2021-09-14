package base

type ListResponse struct {
	Total    int         `json:"total"`    // 总条数
	Page     int         `json:"page"`     // 页数
	PageSize int         `json:"pageSize"` // 每页条数
	Data     interface{} `json:"data"`     // 数据
}

type Response struct {
	Data interface{} `json:"data"` // 数据
}

type StatResponse struct {
	PeopleNumber  int `json:"people_number"`  // 总人数
	RowNumber     int `json:"row_number"`     // 总条数
	ProjectNumber int `json:"project_number"` // 总项目数量
}
