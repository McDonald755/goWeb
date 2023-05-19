package vo

import "goWeb/db"

type GameVo struct {
	PageVo `json:"pageVo"`
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
}

func (vo GameVo) ToEntity() *db.Game {
	return &db.Game{Id: vo.Id, Name: vo.Name, Logo: vo.Logo}
}
