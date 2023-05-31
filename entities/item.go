package entities

import "time"

type Item struct {
	Id        uint      `json:"id" gorm:"type:int(5);primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(25)"`
	Amount    uint      `json:"amount" gorm:"type:int(10)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
