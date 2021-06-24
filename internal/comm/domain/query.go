package commDomain

type SumRes struct {
	Total int64 `json:"total"`
}

// Filed 查询字段结构体
type Filed struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

type Relate struct {
	Value string
	Func  interface{}
}

// Search 查询参数结构体
type Search struct {
	Fields    []*Filed  `json:"fields"`
	Relations []*Relate `json:"relations"`
	OrderBy   string    `json:"order_by"`
	Sort      string    `json:"sort"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
}
