package internal

import (
	"github.com/jayantodpuji/grocerfy/internal/handlers"
	"github.com/jayantodpuji/grocerfy/internal/middlewares"
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
			AuthService: services.NewAuthService(services.AuthServiceDependency{
				JWTKey: app.JWTKey,
			}),
		}),
	})

	groceryListHandler := handlers.NewGroceryListHandler(handlers.GroceryListHandlerDependency{
		GroceryListService: services.NewGroceryListService(services.GroceryListServiceDependency{
			GroceryListRepository: repositories.NewGroceryListRepository(
				repositories.GroceryListRepositoryDependency{
					DB: app.DB,
				},
			),
		}),
	})

	groceryListItemHandler := handlers.NewGroceryListItemHandler(handlers.GroceryListItemHandlerDependency{
		Service: services.NewGroceryListItemService(repositories.NewGroceryListItemRepository(
			repositories.GroceryListItemRepositoryDependency{
				DB: app.DB,
			},
		)),
	})

	v1 := app.Router.Group("/api/v1")

	userV1 := v1.Group("/users")
	userV1.POST("/register", userHandler.Register)
	userV1.POST("/login", userHandler.Login)

	secured := v1.Group("/secured")
	secured.Use(middlewares.AuthMiddleware(app.JWTKey))

	groceryListV1 := secured.Group("/grocery-lists")
	groceryListV1.POST("/create", groceryListHandler.Create)
	groceryListV1.GET("/index", groceryListHandler.Index)
	groceryListV1.GET("/:id/items", groceryListItemHandler.GetByGroceryList)
}
