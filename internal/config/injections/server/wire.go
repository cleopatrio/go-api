package injections

import (
	"github.com/dock-tech/notes-api/internal/config/connections/aws"
	"github.com/dock-tech/notes-api/internal/config/connections/database"
	"github.com/dock-tech/notes-api/internal/delivery/adapters"
	"github.com/dock-tech/notes-api/internal/delivery/controllers"
	"github.com/dock-tech/notes-api/internal/delivery/servers"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"github.com/dock-tech/notes-api/internal/integration/caches"
	"github.com/dock-tech/notes-api/internal/integration/repositories"
	"github.com/dock-tech/notes-api/internal/integration/secrets"
)

func InitializeServer() (adapters.Server, error) {
	cacheClientSet := database.NewCacheGet()
	cacheClientGet := database.NewCacheSet()
	cache := caches.NewCache(cacheClientSet, cacheClientGet)
	config := aws.NewAws()
	secretClient := aws.NewAwsSecretsManager(config)
	secret := secrets.NewSecret(secretClient)
	db := database.NewDb(cache, secret)
	notesRepository := repositories.NewNote(db)
	usersRepository := repositories.NewUser(db)
	errorHandler := controllers.NewErrorHandler()
	usersController := controllers.NewUsersController(
		usecases.CreateUserUseCase(usersRepository),
		usecases.DeleteUserUseCase(usersRepository),
		usecases.GetUserUseCase(usersRepository),
		usecases.ListUsersUseCase(usersRepository),
		errorHandler,
	)
	notesController := controllers.NewNotesController(
		usecases.CreateNoteUseCase(notesRepository),
		usecases.DeleteNoteUseCase(notesRepository),
		usecases.GetNoteUseCase(notesRepository),
		usecases.ListNotesUseCase(notesRepository),
		errorHandler,
	)
	server := servers.NewServer(usersController, notesController)
	return server, nil
}
