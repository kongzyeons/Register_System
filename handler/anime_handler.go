package handler

import (
	"fmt"
	"go_test/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type animeHandler struct {
	animeSrc service.AnimeService
}

func NewAnimeHandler(animeSrc service.AnimeService) AnimeHandler {
	return &animeHandler{animeSrc: animeSrc}

}

func (h *animeHandler) CreateAnime_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)

	user_token := fmt.Sprintf("%v", c.Locals("username"))
	fmt.Println(user_token)
	if user_token != "admin" {
		err := fmt.Errorf("Username token must be of admin ")
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	anime := service.AnimeRequest{}
	if err := c.BodyParser(&anime); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&anime); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}

	animeCreate, err := h.animeSrc.CreateAnime(anime)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	return c.JSON(http.StatusCreated, MessageResponse{
		Status:  true,
		Message: "Create data anime",
		Data:    &fiber.Map{"data": animeCreate}})

}

func (h *animeHandler) GetAllAnime_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
	animes, err := h.animeSrc.GetAllAnime()
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(
		MessageResponse{
			Status:  true,
			Message: "Get data anime",
			Data:    &fiber.Map{"count user": len(animes), "data": animes}})
}

func (h *animeHandler) DeleteAnime_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)

	user_token := fmt.Sprintf("%v", c.Locals("username"))
	fmt.Println(user_token)
	if user_token != "admin" {
		err := fmt.Errorf("Username token must be of admin ")
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	anime_id := c.ParamsID("anime_id")
	anime, err := h.animeSrc.DeleteAnime(anime_id)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(
		MessageResponse{
			Status:  true,
			Message: "Delete data anime",
			Data:    &fiber.Map{"data": anime}})
}
