package mongodb

import (
	"context"
	"time"

	"github.com/gh0stl1m/notification-service/configs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnection() *mongo.Client {

	env := configs.ReadMongoDBConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.URL))
	if err != nil {

		log.Fatalf("Failed to connect to mongoDB: %s", err)
	}

	log.Infof("Mongo connection oppened to %s", env.URL)

	return client
}
