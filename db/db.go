package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoClient(config *viper.Viper) (*mongo.Client, error) {
	connectionURI := config.GetString("mongodb.database_uri")

	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, errConnect := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if errConnect != nil {
		log.Println(errConnect)
		return nil, errConnect
	}

	// Check the connection
	errPing := client.Ping(ctx, nil)
	if errPing != nil {
		fmt.Println("InitMongoClient-err", errPing)
		return nil, errPing
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
