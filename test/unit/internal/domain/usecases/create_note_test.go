package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/dock-tech/notes-api/internal/domain/entities"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"github.com/dock-tech/notes-api/test/mocks"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func BenchmarkCreateNote(b *testing.B) {
	var (
		ctrl             = gomock.NewController(b)
		noteQueueMocked  = mocks.NewMockNoteQueue(ctrl)
		createNoteMocked = mocks.NewMockCreateNoteRepository(ctrl)
		createNote       = usecases.NewCreateNoteUseCase(createNoteMocked, noteQueueMocked)
		ctx              = context.TODO()
		noteEntity       = entities.Note{
			Title:   "Title",
			Content: "Content",
		}
		expectedCreatedNote = entities.Note{
			Id:        uuid.NewString(),
			Title:     "Title",
			Content:   "Content",
			CreatedAt: ptr.Time(time.Now()),
			UpdatedAt: ptr.Time(time.Now()),
		}
	)

	b.Run("CreateNote", func(b *testing.B) {
		createNoteMocked.EXPECT().Create(ctx, noteEntity).Return(&expectedCreatedNote, nil).AnyTimes()
		noteQueueMocked.EXPECT().Publish(ctx, expectedCreatedNote).AnyTimes()

		for i := 0; i < b.N; i++ {
			_, _ = createNote.Create(ctx, noteEntity)
		}
	})

}
