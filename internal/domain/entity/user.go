package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (User) BeforeCreate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("created_at", time.Now())
	db.Statement.SetColumn("updated_at", time.Now())
	db.Statement.SetColumn("id", uuid.New())
	return
}
