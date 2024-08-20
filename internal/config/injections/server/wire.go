package injections

import (
	"github.com/dock-tech/notes-api/internal/delivery/controllers"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"github.com/dock-tech/notes-api/internal/domain/usecases"
	"github.com/dock-tech/notes-api/internal/integration/servers"
	"github.com/google/wire"
)

func InitializeServer() (interfaces.Server, error) {
	wire.Build(

        usecases.NewUsecase,
        controllers.NewController,
        servers.NewServer,
    )
	return nil, nil
}
