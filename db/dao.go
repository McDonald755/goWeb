package db

import (
	log "github.com/sirupsen/logrus"
	"goWeb/config"
	"time"
)

func SaveOrUpdateGame(game Game) (error, int64) {
	var data Game

	//构建基础sql语句
	sql := config.DB.Table("game").Where("game.delete = 0 and game.id = ?", game.Id)
	result := sql.Find(&data)

	if result.RowsAffected == 0 {
		game.CreatedTime = time.Now()
		game.UpdatedTime = time.Now()

		if res := config.DB.Create(&game); res.Error != nil {
			log.Error(res.Error)
			return res.Error, -1
		}

	} else {
		//构建存放修改字段的map
		newData := map[string]interface{}{"updated_time": time.Now()}
		if game.Name != "" {
			newData["name"] = game.Name
		}
		if game.Logo != "" {
			newData["logo"] = game.Logo
		}

		if res := sql.Updates(newData); res.Error != nil {
			log.Error(res.Error)
			return res.Error, -1
		}

	}
	return nil, data.Id
}

func DeleteGame(id string) error {
	// 做逻辑删除
	res := config.DB.Table("game").Where("game.id = ?", id).Update("deleted", "-1")
	if res.Error != nil {
		log.Error(res.Error)
		return res.Error
	}
	return nil
}

func GetGames(game Game) *[]Game {
	var result []Game

	sql := config.DB.Table("game").Where("deleted = 0")

	if game.Name != "" {
		sql.Where("game.name like %?%", game.Name)
	}

	if game.Logo != "" {
		sql.Where("game.logo like %?%", game.Logo)
	}

	if res := sql.Find(&result); res.Error != nil {
		log.Error(res.Error)
		return nil
	}
	return &result
}

func GetGameDetail(id string) *Game {
	var result Game

	res := config.DB.Table("game").Where("deleted = 0 and id = ?", id).First(&result)

	if res.Error != nil {
		log.Error(res.Error)
		return nil
	}
	return &result
}
