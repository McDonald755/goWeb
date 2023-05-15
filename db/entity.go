package db

import "time"

type Game struct {
	Id          int64     `gorm:"id" json:"id"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updatedTime"`
	Delete      int       `gorm:"delete" json:"delete"`
	Name        string    `gorm:"column:name" json:"name"`
	Logo        string    `gorm:"column:logo" json:"logo"`
}

// 制定数据库名
func (Game) TableName() string {
	return "game"
}
