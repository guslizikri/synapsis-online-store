package main

import (
	"log"
	"synapsis-online-store/apps/routers"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	filename := "cmd/config.yaml"
	timeout := 10 * time.Second
	if err := config.LoadConfig(filename); err != nil {
		log.Fatalf("failed to load config from file %s: %v", filename, err)
	}

	db, err := pkg.InitPostgres(config.Cfg.DB)
	if err != nil {
		log.Fatalf("failed to initialize PostgreSQL: %v", err)
	}

	// jika error coba tambahin nama db di uri nya
	uri := "mongodb://root:example@localhost:27017/"
	mongoClient, err := pkg.InitMongoDB(timeout, uri)
	mongoDB := mongoClient.Database("online-store")
	if err != nil || mongoClient == nil {
		log.Fatalf("failed to initialize MongoDB: %v", err)

	}
	redisCLient, err := pkg.InitRedis(timeout)
	if err != nil {
		log.Fatalf("failed to initialize Redis: %v", err)
	}

	router := fiber.New(fiber.Config{
		Prefork: false,
		AppName: config.Cfg.App.Name,
	})

	routers.InitUser(router, db, redisCLient)
	routers.InitProduct(router, db, redisCLient)
	routers.InitCart(router, db, redisCLient)
	routers.InitTransaction(router, db, redisCLient)
	routers.InitReview(router, mongoDB, db, redisCLient)

	if err := router.Listen(config.Cfg.App.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
