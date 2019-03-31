package db

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	config "github.com/spf13/viper"
)

// Ctx is a utility function to get a context
func Ctx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

// GetDatabase returns mongoDb database ref. Should be called once in
// program and should be injected into other APIs
func GetDatabase() *mongo.Database {
	return GetDatabaseWithParams(config.GetString("dbURL"),
		config.GetString("database"))
}

// GetDatabaseWithParams returns mongoDb database ref. Should be called once in
// program and should be injected into other APIs
func GetDatabaseWithParams(dbUrl, database string) *mongo.Database {
	client, err := mongo.Connect(Ctx(), dbUrl)

	if err != nil {
		log.Panic(err)
	}
	// Check the connection
	err = client.Ping(Ctx(), nil)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Connected to MongoDB!")
	return client.Database(database)
}
