package mongodb

import (
	"context"
	"os"

	"github.com/hemilioaraujo/first-go-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URI           = "MONGODB_URI"
	MONGODB_USER_DATABASE = "MONGODB_USER_DATABASE"
)

func NewMongoDBConnection() (*mongo.Database, error) {
	logger.Info("-------- MongoDB connecting --------")

	mongodb_uri := os.Getenv(MONGODB_URI)
	mongodb_database := os.Getenv(MONGODB_USER_DATABASE)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	logger.Info("-------- MongoDB connected --------")
	return client.Database(mongodb_database), nil
}
