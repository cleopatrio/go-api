package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/dock-tech/notes-api/internal/domain/entity"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/dock-tech/notes-api/internal/integration/adapters"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type noteRepository struct {
	connection *gorm.DB
}

func (n *noteRepository) Get(ctx context.Context, userId string, noteId string) (note *entity.Note, err error) {
	err = n.connection.WithContext(ctx).Where(
		&entity.Note{
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

func (n *noteRepository) Create(ctx context.Context, note entity.Note) (createdNote *entity.Note, err error) {
	createdNote = &note
	createdNote.Id = uuid.NewString()
	now := time.Now()
	createdNote.CreatedAt = &now
	createdNote.UpdatedAt = &now
	err = n.connection.WithContext(ctx).Create(&note).Error
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

func (n *noteRepository) Delete(ctx context.Context, userId, noteId string) (err error) {
	tx := n.connection.WithContext(ctx).Delete(&entity.Note{Id: noteId, UserId: userId})
	err = tx.Error
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to delete note with id %s and userId %s", noteId, userId), err.Error())
	}

	if tx.RowsAffected == 0 {
		err = exceptions.NewNotFoundError(fmt.Sprintf("note with id %s and user %s not found", noteId, userId))
		return
	}
	return
}

func (n *noteRepository) List(ctx context.Context, userId string) (notes []*entity.Note, err error) {
	err = n.connection.WithContext(ctx).Where(&entity.Note{UserId: userId}).Find(&notes).Error
	if err != nil {
		err = exceptions.NewInternalServerError(fmt.Sprintf("failed to list notes from userId %s", userId), err.Error())
	}
	return
}

func NewNote(connection *gorm.DB) adapters.NoteRepository {
	return &noteRepository{connection: connection}
}
