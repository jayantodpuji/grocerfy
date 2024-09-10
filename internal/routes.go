package internal

import (
	"github.com/jayantodpuji/grocerfy/internal/handlers"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/services"
)

func Routes(app *Application) {
	userHandler := handlers.NewUserHandler(handlers.UserHandlerDependency{
		UserService: services.NewUserService(services.UserServiceDependency{
			UserRepository: repositories.NewUserRepository(
				repositories.UserRepositoryDependency{
					DB: app.DB,
				},
			),
		}),
	})

	v1 := app.Router.Group("/api/v1")

	userV1 := v1.Group("/users")
	userV1.POST("/register", userHandler.Register)
}
