package service

import "goWeb/db"

func SaveOrUpdateGame(game db.Game) (error, int64) {
	//todo 结构体转换
	return db.SaveOrUpdateGame(game)
}

func DeleteGame(id string) error {
	return db.DeleteGame(id)
}

func GetGames(game db.Game) *[]db.Game {
	//todo 结构体转换
	return db.GetGames(game)
}
func GetGameDetail(id string) *db.Game {
	return db.GetGameDetail(id)
}
