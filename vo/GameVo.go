package vo

type GameVo struct {
	PageVo `json:"pageVo"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
}
