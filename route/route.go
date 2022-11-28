package route

import (
	"go_test/config"
	"go_test/handler"
	"go_test/middleware"
	"go_test/repository"
	"go_test/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouters(route *config.Fiber_fw, connectionDB *mongo.Client) {
	userRepository := repository.NewUserRepositoryDB(connectionDB)
	userService := service.NewUserService(userRepository)
	userHandle := handler.NewUserHandle(userService)

	route.Post("/api/v1/authenticate", userHandle.LoginUser_api)

	private := route.Group("/", middleware.AuthorizationMiddleware)
	private.Post("/api/v1/data", userHandle.CreateUser_api)
	private.Get("/api/v1/data", userHandle.GetUsers_api)
	private.Get("/api/v1/data/:user_id", userHandle.GetUser_api)
	private.Put("/api/v1/data/:user_id", userHandle.UpdateUser_api)
	private.Delete("/api/v1/data/:user_id", userHandle.DeleteUser_api)
}
