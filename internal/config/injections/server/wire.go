package injections

import (
	"github.com/dock-tech/notes-api/internal/config/connections/aws"
	"github.com/dock-tech/notes-api/internal/config/connections/database"
	"github.com/dock-tech/notes-api/internal/delivery/adapters"
	"github.com/dock-tech/notes-api/internal/delivery/controllers"
	"github.com/dock-tech/notes-api/internal/delivery/servers"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"github.com/dock-tech/notes-api/internal/integration/caches"
	"github.com/dock-tech/notes-api/internal/integration/queues"
	"github.com/dock-tech/notes-api/internal/integration/repositories"
	"github.com/dock-tech/notes-api/internal/integration/secrets"
	"gorm.io/gorm"
	"sync"
)

type wire struct {
	Db  *gorm.DB
	Sqs queues.SqsClient
}

var wireInit sync.Once
var wireInstance *wire

func Wire() *wire {
	if wireInstance == nil {
		wireInit.Do(
			func() {
				wireInstance = &wire{}
			},
		)
	}

	return wireInstance
}

func (w *wire) InitializeServer() (adapters.Server, error) {
	config := aws.NewAws()
	if w.Db == nil {
		cacheClientSet := database.NewCacheSet()
		cacheClientGet := database.NewCacheGet()
		cache := caches.NewCache(cacheClientGet, cacheClientSet)
		secretClient := aws.NewAwsSecretsManager(config)
		secret := secrets.NewSecret(secretClient)
		w.Db = database.NewDb(cache, secret)
	}
	notesRepository := repositories.NewNote(w.Db)
	usersRepository := repositories.NewUser(w.Db)
	if w.Sqs == nil {
		w.Sqs = aws.NewAwsSqs(config)
	}
	notesQueue := queues.NewNotesQueue(w.Sqs)
	errorHandler := controllers.NewErrorHandler()
	usersController := controllers.NewUsersController(
		usecases.CreateUserUseCase(usersRepository),
		usecases.DeleteUserUseCase(usersRepository),
		usecases.GetUserUseCase(usersRepository),
		usecases.ListUsersUseCase(usersRepository),
		errorHandler,
	)
	notesController := controllers.NewNotesController(
		*usecases.NewCreateNoteUseCase(notesRepository, notesQueue),
		usecases.DeleteNoteUseCase(notesRepository),
		usecases.GetNoteUseCase(notesRepository),
		usecases.ListNotesUseCase(notesRepository),
		errorHandler,
	)
	server := servers.NewServer(usersController, notesController)
	return server, nil
}
