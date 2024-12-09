package main

import (
	"fmt"
	"log"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
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
}
