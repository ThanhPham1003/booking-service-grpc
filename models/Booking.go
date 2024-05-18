package models

import "time"

type Booking struct {
	StartTime   time.Time `bson:"startTime"`
	EndTime     time.Time `bson:"endTime"`
	Area        float32   `bson:"area"`
	Price       float32   `bson:"price"`
	Note        string    `bson:"note"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
	IsConfirmed bool      `bson:"isConfirmed"`
}
