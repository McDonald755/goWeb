package vo

type PageVo struct {
	PageNum   int64 `json:"pageNum", validate:"min=10,max=0"`
	PageSize  int64 `json:"pageSize"`
	TotalSize int64 `json:"totalSize"`
}

type ResponsePageVo struct {
	PageNum   int64       `json:"pageNum"`
	PageSize  int64       `json:"pageSize"`
	TotalSize int64       `json:"totalSize"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type ResponseVo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
