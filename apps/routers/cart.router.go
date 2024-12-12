package routers

import (
	"synapsis-online-store/apps/handlers"
	"synapsis-online-store/apps/middleware"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func InitCart(router fiber.Router, db *sqlx.DB) {
	repo := repository.NewRepoCart(db)
	svc := services.NewServiceCart(repo)
	handler := handlers.NewHandlerCart(svc)

	cartRouter := router.Group("cart")
	{
		cartRouter.Get("",
			middleware.CheckAuth(),
			handler.GetListCartItem,
		)
		cartRouter.Post("",
			middleware.CheckAuth(),
			handler.CreateCartItem,
		)
		cartRouter.Delete("/:product_id",
			middleware.CheckAuth(),
			handler.DeleteCartItem,
		)
	}
}
