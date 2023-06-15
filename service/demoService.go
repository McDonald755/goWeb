package service

import (
	"goWeb/db"
	"goWeb/vo"
)

func SaveOrUpdateGame(vo vo.GameVo) (error, int64) {
	return db.SaveOrUpdateGame(*gameToEntity(vo))
}

func DeleteGame(id string) error {
	return db.DeleteGame(id)
}

func GetGames(game vo.GameVo) *[]db.Game {
	//结构体转换
	return db.GetGames(*gameToEntity(game), game.PageSize, game.PageNum)
}

func GetGameDetail(id string) *vo.GameVo {
	return db.GetGameDetail(id).ToVo()
}

func gameToEntity(vo vo.GameVo) *db.Game {
	return &db.Game{Id: vo.Id, Name: vo.Name, Logo: vo.Logo}
}
