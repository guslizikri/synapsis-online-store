package main

import (
	"fmt"
	"log"
	"synapsis-online-store/apps/routers"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/config.yaml"
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

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	routers.InitUser(router, db)
	routers.InitProduct(router, db)

	router.Listen(config.Cfg.App.Port)
}
