package database


import (
	"context"
	"log"
	"time"

	"github.com/BramantyoAW/go-graphql-inviteapps/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	
)


type DB struct{
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

func (db *DB) save(input *model.NewUser) *model.User {
	collection := db.client.Database("users").Collection("user")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.User{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: input.Name,
		Email: input.Email,
		Address: input.Address,
		Password: input.Password,
		phone: input.Phone,
	}
}


func (db *DB) findByID(ID string) *model.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("users").Collection("user")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res  := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	user := model.Uses{}
	res.Decode(&user)
	return &user
}


func (db *DB)  all() []*model.User {
	collection := db.client.Database("users").Collection("user")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var users []*model.User
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users
}