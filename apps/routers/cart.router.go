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

	productRouter := router.Group("cart")
	{
		productRouter.Get("",
			middleware.CheckAuth(),
			handler.GetListCartItem,
		)
		productRouter.Post("",
			middleware.CheckAuth(),
			handler.CreateCartItem,
		)
		productRouter.Delete("/:product_id",
			middleware.CheckAuth(),
			handler.DeleteCartItem,
		)
	}
}
