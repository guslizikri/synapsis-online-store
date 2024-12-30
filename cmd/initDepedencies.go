package main

import (
	"context"
	"log"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitPostgres(cfgDB config.DBConfig) (db *sqlx.DB, err error) {
	db, err = pkg.ConnectPostgres(cfgDB)
	if err != nil {
		log.Fatal("postgres not connected with error", err)
	}
	if db == nil {
		log.Fatal("postgres not connected with unknown error")
	}
	log.Println("DB Connected")

	return
}

func InitMongoDB(timeout time.Duration, uri string) (mongoClient *mongo.Client, err error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	mongoClient, err = pkg.ConnectMongo(ctx, uri)
	if err != nil {
		log.Println("mongodb not connected with error", err.Error())
		return
	}

	if mongoClient == nil {
		log.Println("mongodb not connected with unknown error")
		return
	}
	log.Println("mongo connected")
	return
}

func InitRedis(timeout time.Duration) (client *redis.Client, err error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	client, err = pkg.ConnectRedis(ctx, "localhost:6377", "")
	if err != nil {
		log.Println("redis not connected with error", err.Error())
		return
	}

	if client == nil {
		log.Println("redis not connected with unknown error")
		return
	}
	log.Println("redis connected")
	return
}
