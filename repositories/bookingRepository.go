package repositories

import (
	"booking-service-grpc/config/logger"
	"booking-service-grpc/infrastructure"
	"booking-service-grpc/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBookingRepository interface{
	GetAll() ([]models.Booking, error)
}

type BookingRepositoryImpl struct {
	collection *mongo.Collection
}

func NewBookingRepository(db *infrastructure.MongoDB) IBookingRepository {
	collection := db.Client().Database(db.DatabaseName()).Collection("booking")
	return &BookingRepositoryImpl{collection: collection}
}

func (r *BookingRepositoryImpl) GetAll() ([]models.Booking, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	l := logger.NewLogger()
	l.Info("aaaaa")
	if err != nil {
		l.Error("error at fetching booking from db")
		return nil, err
	}
	defer cursor.Close(context.Background())

	var booking []models.Booking
	if err := cursor.All(context.Background(), &booking); err != nil {
		l.Error("error at fetching booking from db")
		return nil, err
	}

	return booking, nil
}