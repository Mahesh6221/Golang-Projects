package database

import (
	"BACKEND/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

// InitializeMongoDB initializes the MongoDB client and returns it
func InitializeMongoDB() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoDB := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongoDB!")
	return client
}

var Client *mongo.Client = InitializeMongoDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("UserDatabase").Collection(collectionName)
	return collection
}

func SaveAssetDetail(client *mongo.Client, assetDetail model.AssetDetail) error {
	collection := client.Database("UserDatabase").Collection("AssetDetails")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, assetDetail)
	return err
}

func GetAssetDetailByEmpID(client *mongo.Client, empID string) (model.AssetDetail, error) {
	var assetDetail model.AssetDetail
	collection := client.Database("UserDatabase").Collection("AssetDetails")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"empID": empID}).Decode(&assetDetail)
	return assetDetail, err
}
