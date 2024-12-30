package routers

import (
	"synapsis-online-store/apps/handlers"
	"synapsis-online-store/apps/middleware"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitReview(router fiber.Router, mongoDB *mongo.Database, db *sqlx.DB, client *redis.Client) {
	repo := repository.NewRepoReview(mongoDB, db)
	svc := services.NewServiceReview(repo)
	handler := handlers.NewHandlerReview(svc)

	productRouter := router.Group("review")
	{
		productRouter.Get("/:product_id", handler.GetReviews)
		productRouter.Post("",
			middleware.CheckAuth(client),
			handler.AddReview,
		)
	}
}
