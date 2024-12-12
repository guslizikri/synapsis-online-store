package routers

import (
	"synapsis-online-store/apps/handlers"
	"synapsis-online-store/apps/middleware"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func InitTransaction(router fiber.Router, db *sqlx.DB) {
	repo := repository.NewRepoTransaction(db)
	svc := services.NewServiceTransaction(repo)
	handler := handlers.NewHandlerTransaction(svc)

	transactionRouter := router.Group("transaction")
	{
		transactionRouter.Post("",
			middleware.CheckAuth(),
			handler.CreateTransaction,
		)
	}
}
