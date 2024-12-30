package pkg

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(ctx context.Context, uri string) (client *mongo.Client, err error) {
	opts := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(ctx, opts)

	if err != nil {
		return
	}

	if err = client.Ping(ctx, nil); err != nil {
		return
	}
	return
}
