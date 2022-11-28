package middleware

import (
	"fmt"
	"go_test/handler"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtSecret = "my_secret_key"
)

func AuthorizationMiddleware(c *fiber.Ctx) error {
	s := c.Get("Authorization")
	tokenString := strings.TrimPrefix(s, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			handler.MessageResponse{
				Status:  false,
				Message: "5555",
				Data:    &fiber.Map{"data": err.Error()}})
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("user_id", claims["user_id"])
		c.Locals("username", claims["username"])
	}

	return c.Next()
}
