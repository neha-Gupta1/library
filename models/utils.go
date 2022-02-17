package models

import (
	"context"
	"log"
	"time"

	"github.com/library/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Database, *mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("error connecting to DB: ", err)
	}
	database := client.Database(utils.Books)
	return database, client, ctx, cancel, err
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
