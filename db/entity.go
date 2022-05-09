package db

import "time"

type PICTURE struct {
	Id          int64     `gorm:"id" json:"id"`
	CreatedTime time.Time `gorm:"created_time" json:"createdTime"`
	UpdatedTime time.Time `gorm:"updated_time" json:"updatedTime"`
	Type        string    `gorm:"type" json:"type"`
	Url         string    `gorm:"url" json:"url"`
	Status      string    `gorm:"status" json:"status"`
	Deleted     int64     `gorm:"deleted" json:"deleted"`
}

func (PICTURE) TableName() string {
	return "PICTURE"
}
