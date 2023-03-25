package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id"`
	Content    string
	CreatedAt  time.Time
	DeletedAt  time.Time
	Properties map[string]interface{}
	Status     int8
}

const (
	PostStatus_Active int8 = iota
	PostStatus_Passive
)
