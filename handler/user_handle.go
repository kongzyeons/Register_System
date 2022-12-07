package handler

import (
	"fmt"
	"go_test/config"
	"go_test/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandle struct {
	userSrc service.UserService
}

func NewUserHandle(userSrc service.UserService) UserHandle {
	return &userHandle{userSrc: userSrc}
}

func connectionFW(c *fiber.Ctx) *config.FiberCtx {
	context := config.NewFiberCtx(c)
	return context
}

func (h *userHandle) CreateUser_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
	user := service.UserRequest{}
	if err := c.BodyParser(&user); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	userCreate, err := h.userSrc.CreateUser(&user)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	// log.Println("success", c.Status(http.StatusCreated))
	return c.JSON(http.StatusBadRequest, MessageResponse{
		Status:  true,
		Message: "Create data",
		Data:    &fiber.Map{"data": userCreate}})

}

func (h *userHandle) LoginUser_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
	user := service.UserLogin{}
	if err := c.BodyParser(&user); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	tokenString, err := h.userSrc.LoginUser(&user)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusCreated).JSON(
		MessageResponse{
			Status:  true,
			Message: "Authenticate",
			Data:    &fiber.Map{"token": tokenString}})

}

<<<<<<< HEAD
func (h *userHandle) GetUsers_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
=======
func (h *userHandle) GetUsers_api(c *fiber.Ctx) error {
>>>>>>> 87a03f1a9861417ff25cd1052f53f1f461cbd4f0
	user_token := fmt.Sprintf("%v", c.Locals("username"))

	if user_token != "admin" {
		err := fmt.Errorf("Username token must be of admin ")
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	users, err := h.userSrc.GetUsers()
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(
		MessageResponse{
			Status:  true,
			Message: "Get data table",
			Data:    &fiber.Map{"count user": len(users), "data": users}})
}

func (h *userHandle) GetUser_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
	user_id := c.ParamsID("user_id")
	id_token := fmt.Sprintf("%v", c.Locals("user_id"))
	if user_id != id_token {
		err := fmt.Errorf("user_id or token is incorrect (repeat logon)")
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	user, err := h.userSrc.GetUser(user_id)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(
		MessageResponse{
			Status:  true,
			Message: "Get data table",
			Data:    &fiber.Map{"data": user}})
}

func (h *userHandle) UpdateUser_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
	user_id := c.ParamsID("user_id")
	id_token := fmt.Sprintf("%v", c.Locals("user_id"))
	user_token := fmt.Sprintf("%v", c.Locals("username"))

	if (user_token != "admin") && (user_id != id_token) {
		err := fmt.Errorf("user_id or token is incorrect (repeat logon)")
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	user := service.UpdateRequest{}
	if err := c.BodyParser(&user); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(err_response(err))
	}
	userUpdate, err := h.userSrc.UpdateUser(user_id, &user)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(
		MessageResponse{
			Status:  true,
			Message: "Update data",
			Data:    &fiber.Map{"data": userUpdate}})

}

func (h *userHandle) DeleteUser_api(ctx *fiber.Ctx) error {
	c := connectionFW(ctx)
	user_id := c.ParamsID("user_id")
	// id_token := fmt.Sprintf("%v", c.Locals("user_id"))
	user_token := fmt.Sprintf("%v", c.Locals("username"))

	if user_token != "admin" {
		err := fmt.Errorf("Username token must be of admin ")
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}

	user, err := h.userSrc.DeleteUser(user_id)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(
		MessageResponse{
			Status:  true,
			Message: "Delete data",
			Data:    &fiber.Map{"data": user}})
}
