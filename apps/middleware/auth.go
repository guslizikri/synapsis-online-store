package middleware

import (
	"context"
	"fmt"
	"log"
	"strings"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func CheckAuth(redis *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return pkg.NewResponse(
				pkg.WithError(pkg.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return pkg.NewResponse(
				pkg.WithError(pkg.ErrorUnauthorized),
			).Send(c)
		}

		token := bearer[1]

		// Periksa apakah token di-blacklist
		result, err := redis.Get(context.Background(), token).Result()

		if err != nil && !strings.Contains(err.Error(), "redis: nil") {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Redis error")
		}

		if result == "blacklisted" {
			return pkg.NewResponse(
				pkg.WithError(pkg.ErrorUnauthorized),
				pkg.WithMessage("token has been blacklisted"),
			).Send(c)
		}

		publicId, role, err := pkg.ValidateToken(token, config.Cfg.App.Encryption.JwtSecret)
		if err != nil {
			log.Println(err.Error())
			return pkg.NewResponse(
				pkg.WithError(pkg.ErrorUnauthorized),
				pkg.WithMessage(err.Error()),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}

func CheckRoles(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("ROLE"))

		isExists := false
		for _, authorizedRole := range authorizedRoles {
			if role == authorizedRole {
				isExists = true
				break
			}
		}

		if !isExists {
			return pkg.NewResponse(
				pkg.WithError(pkg.ErrorForbiddenAccess),
			).Send(c)
		}

		return c.Next()
	}
}
