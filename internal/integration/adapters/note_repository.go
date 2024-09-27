package adapters

import "github.com/dock-tech/notes-api/internal/domain/usecases"

type NoteRepository interface {
	usecases.GetNoteUseCase
	usecases.ListNotesUseCase
	usecases.CreateNoteUseCase
	usecases.DeleteNoteUseCase
}
