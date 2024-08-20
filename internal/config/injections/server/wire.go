//go:build wireinject

package injections

import (
	"github.com/dock-tech/notes-api/internal/delivery/controllers"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"github.com/dock-tech/notes-api/internal/integration/caches"
	"github.com/dock-tech/notes-api/internal/integration/connections"
	"github.com/dock-tech/notes-api/internal/integration/repositories"
	"github.com/dock-tech/notes-api/internal/integration/secrets"
	"github.com/dock-tech/notes-api/internal/integration/servers"
	"github.com/google/wire"
)

func InitializeServer() (interfaces.Server, error) {
	wire.Build(
		connections.NewAws,
		connections.NewAwsSecretsManager,
		connections.NewCacheSet,
		connections.NewCacheGet,
		secrets.NewSecret,
		caches.NewCache,
		connections.NewDb,
		repositories.NewNote,
		repositories.NewUser,
		usecases.NewErrorHandler,
		usecases.NewUsecase,
		controllers.NewController,
		servers.NewServer,
	)
	return nil, nil
}
