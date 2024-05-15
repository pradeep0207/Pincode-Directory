package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pincode struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Circle    string             `json:"circle" bson:"circle,omitempty"`
	Region 	  string             `json:"region" bson:"region,omitempty"`
	Division  string             `json:"division" bson:"division,omitempty"`
	Office	  string 			 `json:"office" bson:"office,omitempty"`
	Pincode   	int			 	 `json:"pincode" bson:"pincode,omitempty"`	
	Officetype string 			 `json:"officetype" bson:"officetype,omitempty"`	
	Delivery   string			 `json:"delivery" bson:"delivery,omitempty"`
	District   string 			 `json:"district" bson:"district,omitempty"`
	StateName	string 			 `json:"state" bson:"state,omitempty"`
}
