package services

import (
	"CRUD/server/conn"
	"CRUD/server/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

//Function To Show Users List
func ShowUsersList() []models.UserDetails {
	res := []models.UserDetails{}
	cursor, err := conn.Collection().Find(conn.GetContext(), bson.D{})
	fmt.Println(cursor)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(conn.GetContext(), &res); err != nil {
		log.Fatal(err)
	}
	return res

}

//Function To Add Users
func AddUsers(user models.UserDetails) {
	// user := models.UserDetails{}
	res, err := conn.Collection().InsertOne(conn.GetContext(), user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted document with ID: ", res.InsertedID)
}

func UpdateDocuments(user models.UserDetails) {
	filter := bson.D{{"formno", user.FormNo}}
	update := bson.D{{"$set", bson.D{{"imageUrl", user.ImageUrl}, {"isDocumentUploaded", true},
		{"isVerified", true}}}}

	fmt.Println(filter, update)
	result, err := conn.Collection().UpdateOne(conn.GetContext(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	if result.MatchedCount != 0 {
		log.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
}

//Function To Update Users
func UpdateUsers(user models.UserDetails) {
	filter := bson.D{{"formno", user.FormNo}}
	update := bson.D{{"$set", bson.D{{"firstname", user.FirstName}, {"lastname", user.LastName},
		{"address", user.Address},
		{"bloodgroup", user.BloodGroup},
		{"phonenumber", user.PhoneNumber},
		{"email", user.Email},
		{"feedback", user.Feedback},
		{"isdonated", user.IsDonated}}}}

	fmt.Println(filter, update)
	result, err := conn.Collection().UpdateOne(conn.GetContext(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		log.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
}

func DeleteUsers(user models.UserDetails) {
	filter := bson.D{{"formno", user.FormNo}}
	result, err := conn.Collection().DeleteOne(conn.GetContext(), filter)
	if err != nil {
		log.Fatal(err)
	}
	if result.DeletedCount != 0 {
		log.Printf("Deleted a new document with ID %v\n", result)
	}
}
