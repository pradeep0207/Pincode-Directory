package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"io"
	"bytes" 
	"github.com/gorilla/mux"

	"example/Go-Api-tutorial/database"
	"example/Go-Api-tutorial/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = database.Db()

var CreateData = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var data models.Pincode
	erro := json.NewDecoder(request.Body).Decode(&data)
	if erro != nil {
		fmt.Print(erro)
		return
	}
	collection := client.Database("golang").Collection("pincode")
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Inserted a second document: ", result)
})

var GetData = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	buffer := bytes.NewBuffer(body)
	var value models.Pincode
	res := json.NewDecoder(buffer).Decode(&value); 
	if res != nil {
		fmt.Println("Request", res)
	}
	fmt.Println("Request:", value)
	fmt.Println("Request:", err)
})

var DeletePersonEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	fmt.Println("Request", params)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var data models.Pincode

	collection := client.Database("golang").Collection("pincode")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&data)
	if err != nil {
		fmt.Println("Request", err)
		return
	}
	_, derr := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if derr != nil {
		fmt.Println("Request", response)
		return
	}
	fmt.Println("Request", response)
})
