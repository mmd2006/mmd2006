package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"-" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
