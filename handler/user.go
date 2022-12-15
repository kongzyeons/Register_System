package handler

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandle interface {
	CreateUser_api(c *fiber.Ctx) error
	GetUsers_api(c *fiber.Ctx) error
	GetUser_api(c *fiber.Ctx) error
	UpdateUser_api(c *fiber.Ctx) error
	DeleteUser_api(c *fiber.Ctx) error
	LoginUser_api(c *fiber.Ctx) error
}
