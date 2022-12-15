package handler

import "github.com/gofiber/fiber/v2"

type AnimeHandler interface {
	CreateAnime_api(ctx *fiber.Ctx) error
	GetAllAnime_api(ctx *fiber.Ctx) error
	DeleteAnime_api(ctx *fiber.Ctx) error
}
