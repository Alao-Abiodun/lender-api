package mongo

import (
	"context"
	"time"
	"log"
	"github.com/Alao-Abiodun/lender-api/internal/domain/user"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(dbURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB successfully")
	return client, nil
}

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(mongodb *mongo.Database) *UserRepository {
	return &UserRepository{ userCollection: mongodb.Collection("users") }
}

func (userRepo *UserRepository) CreateUser(context context.Context, user *user.User) error {
	_, err := userRepo.userCollection.InsertOne(context, user)
	return err
}

func (userRepo *UserRepository) GetUserByID(context context.Context, userID string) (*user.User, error) {
	var user user.User
	err := userRepo.userCollection.FindOne(context, map[string]string{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *UserRepository) UpdateUser(context context.Context, user *user.User) error {
	_, err := userRepo.userCollection.UpdateOne(context, map[string]string{"_id": user.ID}, map[string]interface{}{
		"$set": user,
	})
	return err
}

func (userRepo *UserRepository) GetUsers(context context.Context) ([]*user.User, error) {
	cursor, err := userRepo.userCollection.Find(context, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context)

	var users []*user.User
	for cursor.Next(context) {
		var user user.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}