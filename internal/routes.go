package internal

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jayantodpuji/grocerfy/internal/handlers"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/services"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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

	v1 := app.Router.Group("/api/v1")

	userV1 := v1.Group("/users")
	userV1.POST("/register", userHandler.Register)
	userV1.POST("/login", userHandler.Login)

	secured := v1.Group("/secured")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwt.RegisteredClaims)
		},
		SigningKey: []byte(app.JWTKey),
	}
	secured.Use(echojwt.WithConfig(config))

	groceryListV1 := secured.Group("/grocery-lists")
	groceryListV1.GET("/index", groceryListHandler.Index)
	groceryListV1.POST("/create", groceryListHandler.Create)
}
