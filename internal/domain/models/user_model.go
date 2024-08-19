package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id         string     `gorm:"column:id;primaryKey" json:"id"`
	Name       string     `gorm:"column:name;not null;size:255" json:"name" validate:"required, min=3"`
	Created_at *time.Time `gorm:"column:created_at;not null" json:"created_at"`
	Updated_at *time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (n *User) FromJSON(data []byte) error {
	return json.Unmarshal(data, n)
}
