package vo

type pageReqVo struct {
	PageVo `json:"pageVo"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
}
