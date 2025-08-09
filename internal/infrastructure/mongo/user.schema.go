package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Balance   float64            `bson:"balance" json:"balance"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}