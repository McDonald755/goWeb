package vo

type pageReqVo struct {
	PageVo  `json:"pageVo"`
	Address string `json:"address"`
	Status  string `json:"status"`
	Id      string `json:"id"`
}
