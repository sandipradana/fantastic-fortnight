package middleware

import (
	"fantastic-fortnight/backend/internal/service"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(adminService service.AdminService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")

		if authorizationHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		if !strings.Contains(authorizationHeader, "Bearer") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad request",
			})
		}

		authorizationHeader = strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(authorizationHeader, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_SECRET")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token signature",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad request",
			})
		}

		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		c.Locals("user", claims)

		return c.Next()
	}
}
