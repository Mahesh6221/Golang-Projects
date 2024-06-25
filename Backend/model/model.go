package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	EmpID         string             `json:"empID" bson:"empID"`
	FirstName     string             `json:"firstName" bson:"firstName"`
	LastName      string             `json:"lastName" bson:"lastName"`
	Designation   string             `json:"designation" bson:"designation"`
	Gender        string             `json:"gender" bson:"gender"`
	DOB           string             `json:"dob" bson:"dob"`
	Location      string             `json:"location" bson:"location"`
	ContactNo     string             `json:"contactNo" bson:"contactNo"`
	MailID        string             `json:"mailID" bson:"mailID"`
	Password      string             `json:"Password" bson:"password"`
	Role          string             `json:"role" bson:"role"`
	Type          string             `json:"type"`
	Token         string             `json:"token"`
	Refresh_token string             `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}

type AssetDetail struct {
	EmpID  string `json:"empID" bson:"empID"`
	Laptop struct {
		ModelNumber     string `json:"modelNumber" bson:"modelNumber"`
		OperatingSystem string `json:"operatingSystem" bson:"operatingSystem"`
		AssetType       string `json:"assettype" bson:"assetType"`
		ImageBase64     []byte `json:"imageBase64,omitempty" bson:"image_base64,omitempty"`
		ImageName       string `json:"imageName" bson:"imageName"`
	}
	Mouse struct {
		ModelNumber string `json:"modelnumber" bson:"modelNumber"`
		AssetType   string `json:"assettype" bson:"assetType"`
		ImageBase64 []byte `json:"imageBase64,omitempty" bson:"image_base64,omitempty"`
		ImageName   string `json:"imageName" bson:"imageName"`
	}
	Headphones struct {
		ModelNumber string `json:"modelnumber" bson:"modelNumber"`
		AssetType   string `json:"assettype" bson:"assetType"`
		ImageBase64 []byte `json:"imageBase64,omitempty" bson:"image_base64,omitempty"`
		ImageName   string `json:"imageName" bson:"imageName"`
	}
}
