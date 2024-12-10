package routers

import (
	"synapsis-online-store/apps/handlers"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// fiber.Router karena ini berbentuk interface maka tidak perlu pointer
func InitUser(router fiber.Router, db *sqlx.DB) {
	repo := repository.NewRepoUser(db)
	svc := services.NewServiceUser(repo)
	handler := handlers.NewHandlerUser(svc)
	// karena fiber.router adalah interface, jadi gaperlu di kasih pointer.
	// biar bisa memakai router.group
	userRoute := router.Group("users")
	{
		userRoute.Post("register", handler.Register)
		userRoute.Post("login", handler.Login)
	}
}
