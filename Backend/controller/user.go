package controllers

import (
	"BACKEND/database"
	"BACKEND/model"
	"context"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(c *gin.Context) {
	client := database.InitializeMongoDB()
	EmployeeDetailsCollection := database.GetCollection(client, "EmployeeDetails")
	AssetCollection := database.GetCollection(client, "AssetDetails")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var users []model.User
	if foundUser.Type == "admin" {
		cursor, err := EmployeeDetailsCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var user model.User
			if err := cursor.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user data"})
				return
			}
			if user.Type == "admin" {
				continue
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, gin.H{"foundUser": foundUser, "users": users})
		return
	} else if foundUser.Type == "user" {
		empID := foundUser.ID

		filter := bson.M{"empId": empID}
		var asset model.AssetDetail
		err := AssetCollection.FindOne(ctx, filter).Decode(&asset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving asset"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"foundUser": foundUser, "asset": asset})
		return
	}

	// Fetching employee details
	empCursor, err := EmployeeDetailsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}
	defer empCursor.Close(ctx)

	var employeeDetails []model.User
	if err := empCursor.All(ctx, &employeeDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employees"})
		return
	}

	// Fetching asset details
	assetCursor, err := AssetCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving assets"})
		return
	}
	defer assetCursor.Close(ctx)

	var assetDetails []model.AssetDetail
	if err := assetCursor.All(ctx, &assetDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving assets"})
		return
	}

	// Create a slice to store combined employee and asset details
	combinedDetails := make([]map[string]interface{}, 0)

	// Iterate over employee details
	for _, empDetail := range employeeDetails {
		empAssetDetail := make(map[string]interface{})
		empAssetDetail["details"] = empDetail

		// Find the corresponding asset details for the employee
		for _, assetDetail := range assetDetails {
			if assetDetail.EmpID == empDetail.EmpID {
				empAssetDetail["assets"] = assetDetail
				break
			}
		}
		combinedDetails = append(combinedDetails, empAssetDetail)
	}

	// Return the combinedDetails slice as JSON response
	c.JSON(http.StatusOK, gin.H{"data": combinedDetails})
}

// func UploadFile(c *gin.Context) {
// 	uri := "mongodb://localhost:27017"

// 	// Create a new client and connect to the server
// 	clientOptions := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}

// 	// Ensure the connection is closed when the main function finishes
// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
// 		}
// 	}()

// 	// Specify the database and GridFS bucket
// 	db := client.Database("UserDatabase")
// 	bucket, err := gridfs.NewBucket(db)
// 	if err != nil {
// 		log.Fatalf("Failed to create GridFS bucket: %v", err)
// 	}

// 	// Path to the file you want to upload
// 	filePath := "istockphoto-173863591-1024x1024.jpg"

// 	// Open the file
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	// Upload the file to GridFS
// 	fileName := filepath.Base(filePath) // Use the base name of the file for the GridFS filename
// 	uploadStream, err := bucket.OpenUploadStream(fileName)
// 	if err != nil {
// 		log.Fatalf("Failed to open upload stream: %v", err)
// 	}
// 	defer uploadStream.Close()

// 	// Copy the file data into the upload stream
// 	fileSize, err := io.Copy(uploadStream, file)
// 	if err != nil {
// 		log.Fatalf("Failed to write to upload stream: %v", err)
// 	}

// 	fmt.Printf("File %s uploaded to GridFS with size %d bytes\n", fileName, fileSize)
// 	c.String(200, fmt.Sprintf("File %s uploaded to GridFS with size %d bytes\n", fileName, fileSize))
// }

// func UploadFile(c *gin.Context) {
// 	uri := "mongodb://localhost:27017"

// 	// Create a new client and connect to the server
// 	clientOptions := options.Client().ApplyURI(uri)
// 	client, err := mongo.Connect(context.Background(), clientOptions)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MongoDB: %v", err)
// 	}

// 	// Ensure the connection is closed when the main function finishes
// 	defer func() {
// 		if err := client.Disconnect(context.Background()); err != nil {
// 			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
// 		}
// 	}()

// 	// Specify the database and GridFS bucket
// 	db := client.Database("UserDatabase")
// 	bucket, err := gridfs.NewBucket(db)
// 	if err != nil {
// 		log.Fatalf("Failed to create GridFS bucket: %v", err)
// 	}

// 	// Path to the file you want to upload
// 	filePath := "./image"

// 	// Open the file
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	// Upload the file to GridFS
// 	fileName := filepath.Base(filePath)
// 	uploadStream, err := bucket.OpenUploadStream(fileName)
// 	if err != nil {
// 		log.Fatalf("Failed to open upload stream: %v", err)
// 	}
// 	defer uploadStream.Close()

// 	// Copy the file data into the upload stream
// 	fileSize, err := io.Copy(uploadStream, file)
// 	if err != nil {
// 		log.Fatalf("Failed to write to upload stream: %v", err)
// 	}

// 	fmt.Printf("File %s uploaded to GridFS with size %d bytes\n", fileName, fileSize)
// 	c.String(200, fmt.Sprintf("File %s uploaded to GridFS with size %d bytes\n", fileName, fileSize))
// }

func Uploadfile(c *gin.Context) {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer client.Disconnect(context.Background())

	// Specify the directory where the images are located
	imageDir := "./image"

	// Read the list of files in the directory
	files, err := ioutil.ReadDir(imageDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read image directory"})
		return
	}

	// Iterate over the files in the directory
	for _, file := range files {
		// Read the image data into a byte slice
		imageData, err := ioutil.ReadFile(filepath.Join(imageDir, file.Name()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read image content"})
			return
		}

		// Update existing documents with new image URLs
		filter := bson.M{} // Assuming you want to update all documents
		update := bson.M{
			"$set": bson.M{
				"laptop.image_base64":    imageData,
				"mouse.image_base64":     imageData,
				"headphones.image_base64": imageData,
			},
		}

		// Update documents in the collection
		collection := client.Database("UserDatabase").Collection("AssetDetails")
		_, err = collection.UpdateMany(context.Background(), filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset details"})
			return
		}
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Asset details with images updated successfully"})
}

func GetAssetDetails(c *gin.Context) {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer client.Disconnect(context.Background())

	// Define filter to match all documents (assuming you want to fetch all documents)
	filter := bson.M{}

	// Retrieve documents from the collection
	collection := client.Database("UserDatabase").Collection("AssetDetails")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch asset details"})
		return
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor to decode and collect the documents
	var assetDetails []bson.M
	for cursor.Next(context.Background()) {
		var assetDetail bson.M
		if err := cursor.Decode(&assetDetail); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode asset details"})
			return
		}
		assetDetails = append(assetDetails, assetDetail)
	}
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error while fetching asset details"})
		return
	}

	// Return the fetched documents
	c.JSON(http.StatusOK, gin.H{"asset_details": assetDetails})
}