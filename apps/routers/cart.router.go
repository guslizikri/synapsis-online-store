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

func InitCart(router fiber.Router, db *sqlx.DB, client *redis.Client) {
	repo := repository.NewRepoCart(db)
	svc := services.NewServiceCart(repo)
	handler := handlers.NewHandlerCart(svc)

	cartRouter := router.Group("cart")
	{
		cartRouter.Get("",
			middleware.CheckAuth(client),
			handler.GetListCartItem,
		)
		cartRouter.Post("",
			middleware.CheckAuth(client),
			handler.CreateCartItem,
		)
		cartRouter.Delete("/:product_id",
			middleware.CheckAuth(client),
			handler.DeleteCartItem,
		)
	}
}
