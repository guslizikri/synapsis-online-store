package routers

import (
	"synapsis-online-store/apps/handlers"
	"synapsis-online-store/apps/middleware"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func InitProduct(router fiber.Router, db *sqlx.DB, client *redis.Client) {
	repo := repository.NewRepoProduct(db)
	svc := services.NewServiceProduct(repo)
	handler := handlers.NewHandlerProduct(svc)

	productRouter := router.Group("product")
	{
		productRouter.Get("", handler.GetListProduct)
		productRouter.Post("",
			middleware.CheckAuth(client),
			// infrafiber.CheckRoles([]string{string(users.Role_Admin)}),
			handler.CreateProduct,
		)
	}
}
