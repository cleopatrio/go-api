package entity

import (
	"time"
)

type Note struct {
	Id        string     `gorm:"column:id;primaryKey" json:"id"`
	Title     string     `gorm:"column:title;not null" json:"title" validate:"required,min=3"`
	Content   string     `gorm:"column:content;not null" json:"content" validate:"required,min=3"`
	UserId    string     `gorm:"column:user_id;not null" json:"user_id"  validate:"required"`
	User      User       `gorm:"foreignKey:UserId" json:"-" validate:"-"`
	CreatedAt *time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (Note) TableName() string {
	return "notes"
}
