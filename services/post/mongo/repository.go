package mongo

import (
	"context"
	"log"
	"time"

	"github.com/recepdmr/news-feed/services/post/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	c *mongo.Client
}

var _ domain.PostRepository = repository{}

func (r repository) Delete(p domain.Post) error {
	_, err := r.set().DeleteOne(context.TODO(), bson.M{"_id": p.ID})
	return err
}

func (r repository) GetAll() []domain.Post {
	cursor, err := r.set().Find(context.TODO(), bson.D{{Key: "status", Value: domain.PostStatus_Active}})

	if err != nil {
		log.Fatalf("GetAll thrown %v", err)
		return []domain.Post{}
	}

	var posts []domain.Post
	if err = cursor.All(context.TODO(), &posts); err != nil {
		log.Fatalf("GetAll thrown %v", err)
		return []domain.Post{}
	}
	return posts
}

func (r repository) GetById(id string) (*domain.Post, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	var result domain.Post
	err = r.set().FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}

func (r repository) Insert(p domain.Post) error {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()
	_, err := r.set().InsertOne(context.TODO(), p)
	return err
}

func NewMongoDBPostRepository() domain.PostRepository {
	client := CreateClient()
	return repository{c: client}
}

func (r repository) set() *mongo.Collection {

	database := r.c.Database("post")
	return database.Collection("posts")
}
