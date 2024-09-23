package internal

import (
	"net/http"

	"github.com/jayantodpuji/grocerfy/internal/handlers"
	"github.com/jayantodpuji/grocerfy/internal/middlewares"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4/middleware"
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
			GroceryListItemRepository: repositories.NewGroceryListItemRepository(
				repositories.GroceryListItemRepositoryDependency{
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

	app.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	v1 := app.Router.Group("/api/v1")

	userV1 := v1.Group("/users")
	userV1.POST("/register", userHandler.Register)
	userV1.POST("/login", userHandler.Login)

	secured := v1.Group("/secured")
	secured.Use(middlewares.AuthMiddleware(app.JWTKey))

	groceryListV1 := secured.Group("/lists")
	groceryListV1.POST("/", groceryListHandler.Create)
	groceryListV1.GET("/", groceryListHandler.Index)
	groceryListV1.GET("/:id", groceryListHandler.Detail)
	groceryListV1.PATCH("/:id", groceryListHandler.Update)
	groceryListV1.DELETE("/:id", groceryListHandler.Delete)

	groceryListItemV1 := secured.Group("/items")
	groceryListItemV1.POST("/", groceryListItemHandler.Create)
	groceryListItemV1.GET("/:id", groceryListItemHandler.Detail)
	groceryListItemV1.PATCH("/:id", groceryListItemHandler.Update)
	groceryListItemV1.DELETE("/:id", groceryListItemHandler.Delete)
}
