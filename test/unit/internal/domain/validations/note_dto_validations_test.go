package main

import (
	"testing"
	"time"

	"github.com/dock-tech/notes-api/internal/delivery/validations"
	"github.com/dock-tech/notes-api/internal/domain/exceptions"
	"github.com/stretchr/testify/assert"
)

type Note struct {
	Id        string     `json:"id"`
	Title     string     `json:"title" validate:"required,min=3"`
	Content   string     `json:"content" validate:"required,min=3"`
	UserId    string     `json:"user_id" validate:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func TestNoteValidationWithErrors(t *testing.T) {

	tests := []struct {
		name    string
		note    Note
		wantErr bool
		err     error
	}{
		{
			name: "valid note",
			note: Note{
				Id:      "1",
				Title:   "Valid Title",
				Content: "Valid Content",
				UserId:  "user123",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "missing title",
			note: Note{
				Id:      "2",
				Content: "Valid Content",
				UserId:  "user123",
			},
			wantErr: true,
			err:     exceptions.NewValidationError("'title' is required"),
		},
		{
			name: "short title",
			note: Note{
				Id:      "3",
				Title:   "Hi",
				Content: "Valid Content",
				UserId:  "user123",
			},
			wantErr: true,
			err:     exceptions.NewValidationError("'title' should be greater in length"),
		},
		{
			name: "missing content",
			note: Note{
				Id:     "4",
				Title:  "Valid Title",
				UserId: "user123",
			},
			wantErr: true,
			err:     exceptions.NewValidationError("'content' is required"),
		},
		{
			name: "missing user_id",
			note: Note{
				Id:      "5",
				Title:   "Valid Title",
				Content: "Valid Content",
			},
			wantErr: true,
			err:     exceptions.NewValidationError("'user_id' is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validations.Validate(&tt.note)

			if tt.wantErr {
				assert.Equal(t, err, tt.err)
			}

		})
	}
}
