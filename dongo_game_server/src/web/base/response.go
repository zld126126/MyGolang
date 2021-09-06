package base

type Response struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Data     interface{} `json:"data"`
}
