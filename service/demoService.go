package service

import (
	"goWeb/db"
	"goWeb/vo"
)

func SaveOrUpdateGame(vo vo.GameVo) (error, int64) {

	return db.SaveOrUpdateGame(*vo.ToEntity())
}

func DeleteGame(id string) error {
	return db.DeleteGame(id)
}

func GetGames(game vo.GameVo) *[]db.Game {
	//结构体转换
	entity := game.ToEntity()
	return db.GetGames(*entity, game.PageSize, game.PageNum)
}

func GetGameDetail(id string) *vo.GameVo {
	return db.GetGameDetail(id).ToVo()
}
