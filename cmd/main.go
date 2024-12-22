package main

import (
	"context"
	"fmt"
	"log"
	"synapsis-online-store/apps/routers"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	filename := "cmd/config.yaml"
	// timeout := 2 * time.Second
	if err := config.LoadConfig(filename); err != nil {
		log.Fatal("error file config.yaml", err)
	}
	db, err := pkg.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		log.Fatal("error db start", err)
	}
	if db != nil {
		fmt.Println("DB Connected")
	}

	client, err := pkg.ConnectRedis(context.Background(), "localhost:6377", "")
	if err != nil {
		log.Println("redis not connected with error", err.Error())
		return
	}

	if client == nil {
		log.Println("redis not connected with unknown error")
		return
	}
	log.Println("redis connected")

	router := fiber.New(fiber.Config{
		Prefork: false,
		AppName: config.Cfg.App.Name,
	})

	routers.InitUser(router, db, client)
	routers.InitProduct(router, db, client)
	routers.InitCart(router, db, client)
	routers.InitTransaction(router, db, client)

	router.Listen(config.Cfg.App.Port)
}

// func redis(timeout time.Duration) () {
// 	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
// 	defer cancel()

// 	client, err := pkg.ConnectRedis(ctx, "localhost:6377", "")
// 	if err != nil {
// 		log.Println("redis not connected with error", err.Error())
// 		return
// 	}

// 	if client == nil {
// 		log.Println("redis not connected with unknown error")
// 		return
// 	}
// 	log.Println("redis connected")
// }
