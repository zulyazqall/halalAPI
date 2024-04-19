package datasource

import (
	"context"
	"halalapi/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(cfg config.DatabaseConfig) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.Host).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		// panic(err)
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != mongo.ErrNilValue {
		return nil, err
	}

	return client, nil
}
