package models

import (
	"time"
)

type Note struct {
	Id         string     `gorm:"column:id;primaryKey" json:"id"`
	UserId     string     `gorm:"column:user_id;not null" json:"user_id"  validate:"required"`
	Title      string     `gorm:"column:title;not null" json:"title" validate:"required, min=3"`
	Created_at *time.Time `gorm:"column:created_at;not null" json:"created_at"`
	Updated_at *time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (Note) TableName() string {
	return "notes"
}
