package routers

import (
	"synapsis-online-store/apps/handlers"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func InitProduct(router fiber.Router, db *sqlx.DB) {
	repo := repository.NewRepoProduct(db)
	svc := services.NewServiceProduct(repo)
	handler := handlers.NewHandlerProduct(svc)

	productRouter := router.Group("product")
	{
		productRouter.Get("", handler.GetListProduct)
		productRouter.Post("",
			// infrafiber.CheckAuth(),
			// infrafiber.CheckRoles([]string{string(users.Role_Admin)}),
			handler.CreateProduct,
		)
		// productRouter.Get("/sku/:sku", handler.GetDetailProduct)
	}
}
