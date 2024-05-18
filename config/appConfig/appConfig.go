package appConfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server struct {
		Host              string `env:"HOST"`
		Grpc_port         string `env:"GRPC_PORT"`
		Grpc_gateway_port string `env:"GRPC_GATEWAY_PORT"`
	}
	MongoDB struct {
		Mongo_URI string `env:"MONGO_URI"`
		Database  string `env:"DATABASE_NAME"`
	}
	
}

func LoadConfig() (*AppConfig, error) {
	// Load configuration from .env file

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading config: ", err)
		return nil, err
	}

	// Parse configuration
	config := AppConfig{
		Server: struct {
			Host              string `env:"HOST"`
			Grpc_port         string `env:"GRPC_PORT"`
			Grpc_gateway_port string `env:"GRPC_GATEWAY_PORT"`
		}{
			Host:              os.Getenv("SERVER_HOST"),
			Grpc_port:         os.Getenv("GRPC_PORT"),
			Grpc_gateway_port: os.Getenv("GRPC_GATEWAY_PORT"),
		},
		MongoDB: struct {
			Mongo_URI string `env:"MONGO_URI"`
			Database  string `env:"DATABASE_NAME"`
		}{
			Mongo_URI: os.Getenv("MONGO_URI"),
			Database:  os.Getenv("DATABASE_NAME"),
		},
	}

	return &config, nil
}