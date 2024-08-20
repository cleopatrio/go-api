package repositories

import (
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
	"gorm.io/gorm"
)

type note struct {
	connection *gorm.DB
}

func (n note) List(userId string) (notes []*models.Note, err error) {
	err = n.connection.Find(&notes).Error
	return
}

func (n note) Get(userId string, noteId string) (note *models.Note, err error) {
	err = n.connection.Where(
		&models.Note{
			UserId: userId,
			Id:     noteId,
		}).First(&note).Error
	return
}

func (n note) Create(note models.Note) (err error) {
	err = n.connection.Create(note).Error
	return
}

func (n note) Delete(noteId string) (err error) {
	err = n.connection.Delete(&models.Note{}, noteId).Error
	return
}

func (n note) ListNote(userId string) (notes []*models.Note, err error) {
	err = n.connection.Where(&models.Note{UserId: userId}).Find(&notes).Error
	return
}

func NewNote(connection *gorm.DB) interfaces.NoteRepository {
	return &note{connection: connection}
}
