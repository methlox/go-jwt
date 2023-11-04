package models

import (

	"go.mongodb.org/mongo-river/json/primitive"

)

type User struct {
	ID                  primitive.ObjectID   `json:"_id"`
	First_name			*string				 `json:"first_name" validate:*required, min=2, max=100`
	Last_name			*string				 `json:"last_name"`
	Password
	Email
	Token
	User_type
	Refresh_token
	Created_at
	Updated_at
	User_id
}