package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRepository interface {
	Create(User_db) (*User_db, error)
	Login(user User_db) (*User_db, error)
	Getall() ([]User_db, error)
	GetById(string) (*User_db, error)
	Update(user_id string, updatedUser User_db) (*User_db, error)
	Delete(user_id string) (*User_db, error)
}

type User_db struct {
	UserID           primitive.ObjectID `json:"userID" bson:"userID"`
	Username         string             `json:"username" bson:"username" validate:"required"`
	Password         string             `json:"password" bson:"password" validate:"required"`
	Firstname        string             `json:"firstname" bson:"firstname" validate:"required"`
	Lastname         string             `json:"lastname" bson:"lastname" validate:"required"`
	Birthdate        string             `json:"birthdate" bson:"birthdate" validate:"required"`
	Programing_skill []string           `json:"programing_skill" bson:"programing_skill" validate:"required"`
	Version          int                `json:"version" bson:"version"`
}
