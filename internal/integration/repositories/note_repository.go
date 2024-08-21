package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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
	now := time.Now()
	createdNote.CreatedAt = &now
	createdNote.UpdatedAt = &now
	err = n.connection.Create(&note).Error
	if err != nil {
		//@cleopatrio @brienze1 precisamos ajustar o belongs to do GORM. Eu n vou conseguir fazer isso. Conseguem dar uma olhada?
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			//todo: usar errors.Is(err, gorm.ErrForeignKeyViolation)
			err = exceptions.NewNotFoundError(fmt.Sprintf("user with id %s not found", note.UserId))
			return
		}
		err = exceptions.NewInternalServerError("failed to create note", err.Error())
	}
	return
}

func (n note) Delete(ctx context.Context, noteId string) (err error) {
	tx := n.connection.Delete(&models.Note{Id: noteId})
	err = tx.Error
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to delete note with id %s", noteId), err.Error())
	}

	if tx.RowsAffected == 0 {
		err = exceptions.NewNotFoundError(fmt.Sprintf("note with id %s not found", noteId))
		return
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
