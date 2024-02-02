package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClients(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var dbMongoURL string
	var isAuth bool
	if username == "" && password == "" {
		dbMongoURL = fmt.Sprintf("mongodb://%s:%s", host, port) // mongodb:localhost//:port
	} else {
		isAuth = true
		dbMongoURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port) // connectionString , mongo://Postgresql:pass@localhost:port
	}

	optionsClient := options.Client().ApplyURI(dbMongoURL)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		optionsClient.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	//connect
	client, err := mongo.Connect(ctx, optionsClient)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongoDB error: %v", err)
	}

	//ping
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongoDB error: %v", err)
	}

	return client.Database(database), nil

}
