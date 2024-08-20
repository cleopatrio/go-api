package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type note struct {
	connection *gorm.DB
}

func (n note) List(ctx context.Context, userId string) (notes []*models.Note, err error) {
	err = n.connection.Where(&models.Note{UserId: userId}).Find(&notes).Error
	return
}

func (n note) Get(ctx context.Context, userId string, noteId string) (note *models.Note, err error) {
	err = n.connection.Where(
		&models.Note{
			UserId: userId,
			Id:     noteId,
		}).First(&note).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exceptions.NewNotFoundError(fmt.Sprintf("note with id %s not found", noteId))
		}
	}
	return
}

func (n note) Create(ctx context.Context, note models.Note) (createdNote *models.Note, err error) {
	createdNote = &note
	createdNote.Id = uuid.NewString()
	err = n.connection.Create(&note).Error
	if err != nil {
		//todo: what if userId does not exists?
		err = exceptions.NewInternalServerError("failed to create note", err.Error())
	}
	return
}

func (n note) Delete(ctx context.Context, noteId string) (err error) {
	err = n.connection.Delete(&models.Note{Id: noteId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exceptions.NewNotFoundError(fmt.Sprintf("note with id %s not found", noteId))
			return
		}
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to delete note with id %s", noteId), err.Error())
	}
	return
}

func (n note) ListNote(ctx context.Context, userId string) (notes []*models.Note, err error) {
	err = n.connection.Where(&models.Note{UserId: userId}).Find(&notes).Error
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to list notes from userId %s", userId), err.Error())
	}
	return
}

func NewNote(connection *gorm.DB) interfaces.NoteRepository {
	return &note{connection: connection}
}
