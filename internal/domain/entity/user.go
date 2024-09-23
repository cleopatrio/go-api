package entity

import (
	"time"
)

type User struct {
	Id        string     `gorm:"column:id;primaryKey" json:"id"`
	Name      string     `gorm:"column:name;not null;size:255" json:"name" validate:"required,min=3"`
	Note      []*Note    `gorm:"foreignKey:UserId;references:Id" json:"-"`
	CreatedAt *time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
