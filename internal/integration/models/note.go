package models

import (
	"time"

	"github.com/dock-tech/notes-api/internal/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO arrumar entirades (json com dto no delivery) (gorm com model no integration)
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

func (Note) BeforeCreate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("created_at", time.Now())
	db.Statement.SetColumn("updated_at", time.Now())
	db.Statement.SetColumn("id", uuid.New())
	return
}

func (Note) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("updated_at", time.Now())
	return
}

func (n Note) ToEntity() *entities.Note {
	return &entities.Note{
		Id:        n.Id,
		Title:     n.Title,
		Content:   n.Content,
		UserId:    n.UserId,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}

type Notes []*Note

func (n Notes) ToEntities() []*entities.Note {
	notes := make([]*entities.Note, len(n))
	for i, note := range n {
		notes[i] = note.ToEntity()
	}

	return notes
}
