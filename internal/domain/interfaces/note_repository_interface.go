package interfaces

import "github.com/dock-tech/notes-api/internal/domain/models"

type NoteRepository interface {
	Get(userId string, noteId string) (*models.Note, error)
	Create(note models.Note) error
	Delete(noteId string) error
	List(userId string) ([]*models.Note, error)
}
