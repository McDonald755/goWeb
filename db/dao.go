package db

func GetGameList() *[]Game {
	var result []Game
	//
	//res := config.DB.Table("game").Where("type = ? and status = ? and deleted = 0", t, status).Find(&result)
	//if res.Error != nil {
	//	log.Error(res.Error)
	//}
	return &result
}
