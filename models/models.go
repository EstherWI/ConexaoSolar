package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Usuario struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Email  string             `json:"email" bson:"email,omitempty"`
	Senha  string             `json:"senha" bson:"senha,omitempty"`
	Active bool               `bson:"active" json:"active"`
}
