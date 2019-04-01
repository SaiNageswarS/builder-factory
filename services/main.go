package main

import (
	"fmt"
	"log"
	"net"
	"os"

	config "github.com/spf13/viper"
	"google.golang.org/grpc"
	// "github.com/SaiNageswarS/builder-factory/services/db"
)

const (
	port = ":50051"
)

func main() {
	config.SetConfigType("yaml")

	runEnv := os.Getenv("RUN_ENV")
	configPath := fmt.Sprintf("config.%s.yaml", runEnv)

	log.Printf("Loading config from %s file", configPath)
	config.SetConfigFile(configPath)
	config.ReadInConfig()
	log.Println("Loaded config")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// mgoDb := odm.GetDatabaseWithParams(config.GetString("dbURL"),
	// config.GetString("database"))
	// orgCollection := db.NewOrgCollection(mgoDb)

	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
