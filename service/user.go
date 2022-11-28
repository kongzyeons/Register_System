package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(user_id string) (*UserResponse, error)
	CreateUser(*UserRequest) (*UserResponse, error)
	UpdateUser(user_id string, requestUser *UpdateRequest) (*UserResponse, error)
	DeleteUser(user_id string) (*DeleteResponse, error)
	LoginUser(loginUser *UserLogin) (string, error)
}

type UserRequest struct {
	Username         string   `json:"username" bson:"username" validate:"required"`
	Password         string   `json:"password" bson:"password" validate:"required"`
	Firstname        string   `json:"firstname" bson:"firstname" validate:"required"`
	Lastname         string   `json:"lastname" bson:"lastname" validate:"required"`
	Birthdate        string   `json:"birthdate" bson:"birthdate" validate:"required"`
	Programing_skill []string `json:"programing_skill" bson:"programing_skill" validate:"required"`
}

type UserResponse struct {
	UserID           primitive.ObjectID `json:"userID" bson:"userID"`
	Username         string             `json:"username" bson:"username" validate:"required"`
	Firstname        string             `json:"firstname" bson:"firstname" validate:"required"`
	Lastname         string             `json:"lastname" bson:"lastname" validate:"required"`
	Birthdate        string             `json:"birthdate" bson:"birthdate" validate:"required"`
	Programing_skill []string           `json:"programing_skill" bson:"programing_skill" validate:"required"`
}

type UpdateRequest struct {
	Firstname        string   `json:"firstname" bson:"firstname" validate:"required"`
	Lastname         string   `json:"lastname" bson:"lastname" validate:"required"`
	Birthdate        string   `json:"birthdate" bson:"birthdate" validate:"required"`
	Programing_skill []string `json:"programing_skill" bson:"programing_skill" validate:"required"`
}

type DeleteResponse struct {
	UserID   primitive.ObjectID `json:"userID" bson:"userID"`
	Username string             `json:"username" bson:"username" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}
