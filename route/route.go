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
	userHandle := handler.NewUserHandler(userService)
	route.Post("/api/v1/authenticate", userHandle.LoginUser_api)
	private_user := route.Group("/api/v1/data", middleware.AuthorizationMiddleware)
	private_user.Post("/", userHandle.CreateUser_api)
	private_user.Get("/", userHandle.GetUsers_api)
	private_user.Get("/:user_id", userHandle.GetUser_api)
	private_user.Put("/:user_id", userHandle.UpdateUser_api)
	private_user.Delete("/:user_id", userHandle.DeleteUser_api)

	animeRepository := repository.NewAnimeRepositoryDB(connectionDB)
	animeService := service.NewAnimeService(animeRepository)
	animeHandle := handler.NewAnimeHandler(animeService)
	private_anime := route.Group("/api/v1/anime", middleware.AuthorizationMiddleware)
	private_anime.Post("/", animeHandle.CreateAnime_api)
	private_anime.Get("/", animeHandle.GetAllAnime_api)
	private_anime.Delete("/:anime_id", animeHandle.DeleteAnime_api)
}
