package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	db *mongo.Client
}

func NewUserRepositoryDB(db *mongo.Client) UserRepository {
	return &userRepositoryDB{db: db}
}

func (r *userRepositoryDB) getCollection() *mongo.Collection {
	collection := r.db.Database(DBName_user).Collection(collection_user)
	return collection
}

func (r *userRepositoryDB) Create(user User_db) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check username
	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)
	// if found username
	if err == nil {
		return nil, fmt.Errorf("username not ready")
	}

	// create and check userID
	user_id := uuid.New().String()
	err = collection.FindOne(ctx, bson.M{"userID": user_id}).Decode(&user)
	user.UserID = user_id
	for err == nil {
		user_id := uuid.New().String()
		err = collection.FindOne(ctx, bson.M{"userID": user_id}).Decode(&user)
		user.UserID = user_id
	}

	_, err = collection.InsertOne(ctx, &user)

	return &user, err

}

func (r *userRepositoryDB) Login(user User_db) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getuser := User_db{}
	// check username
	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&getuser)
	if err != nil {
		return nil, fmt.Errorf("Not found username")
	}
	// check username and password
	err = collection.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}).Decode(&getuser)
	if err != nil {
		return nil, fmt.Errorf("Invalid Login or password.")
	}
	return &getuser, err

}

func (r *userRepositoryDB) Getall() ([]User_db, error) {
	users := []User_db{}
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := collection.Find(ctx, bson.M{})
	defer results.Close(ctx)
	for results.Next(ctx) {
		singleUser := User_db{}
		err = results.Decode(&singleUser)
		users = append(users, singleUser)
	}
	if err != nil {
		return nil, err
	}
	return users, err
}

func (r *userRepositoryDB) GetById(user_id string) (*User_db, error) {
	user := User_db{}
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"userID": user_id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepositoryDB) Update(user_id string, updatedUser User_db) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// check verion update
	user_db, err := r.GetById(user_id)
	if err != nil {
		return nil, err
	}
	version := user_db.Version
	for updatedUser.Version != version {
		user_db, _ = r.GetById(user_id)
		updatedUser.Version = user_db.Version
	}
	updatedUser.Version = updatedUser.Version + 1

	update := bson.M{"firstname": updatedUser.Firstname, "lastname": updatedUser.Lastname,
		"birthdate": updatedUser.Birthdate, "programing_skill": updatedUser.Programing_skill,
		"version": updatedUser.Version}
	result, err := collection.UpdateOne(ctx, bson.M{"userID": user_id}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 1 {
		err := collection.FindOne(ctx, bson.M{"userID": user_id}).Decode(&updatedUser)
		if err != nil {
			return nil, err
		}
		return &updatedUser, err
	}
	return &updatedUser, err

}

func (r *userRepositoryDB) Delete(user_id string) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// delete from userid
	user_db, err := r.GetById(user_id)
	if err != nil {
		return nil, err
	}
	result, err := collection.DeleteOne(ctx, bson.M{"userID": user_id})
	if err != nil {
		return nil, err
	}
	if result.DeletedCount < 1 {
		return nil, fmt.Errorf("Not found User_id")
	}
	return user_db, err

}
