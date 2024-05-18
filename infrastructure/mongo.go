package infrastructure

import (
	appConfig "booking-service-grpc/config/appConfig"
	"booking-service-grpc/config/logger"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	dbName string
}


func NewMongoDB(cfg *appConfig.AppConfig) (m *MongoDB, err error) {
	l := logger.NewLogger()
	clientOptions := options.Client().ApplyURI(cfg.MongoDB.Mongo_URI)
	l.Info(cfg.MongoDB.Mongo_URI)
	
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// log.Fatal(err)
		l.Error("error connecting to MongoDB")
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// log.Fatal("Error when ping:", err)
		l.Error("error when ping:")
	}
	m = &MongoDB{
		client: client,
		dbName: cfg.MongoDB.Database,
	}
	return m, err
}

func (db *MongoDB) Client() *mongo.Client {
	return db.client
}

func (db *MongoDB) DatabaseName() string {
	return db.dbName
}
