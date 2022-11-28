package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	db *mongo.Client
}

func NewUserRepositoryDB(db *mongo.Client) UserRepository {
	return &userRepositoryDB{db: db}
}

const (
	DBName     = "" //kong_test
	collection = "" //user2
)

func (r *userRepositoryDB) getCollection() *mongo.Collection {
	collection := r.db.Database(DBName).Collection(collection)
	return collection
}

func (r *userRepositoryDB) Create(user User_db) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)
	// if found username
	if err == nil {
		return nil, fmt.Errorf("This username : '%s' already exists.", user.Username)
	}

	_, err = collection.InsertOne(ctx, &user)

	return &user, err

}

func (r *userRepositoryDB) Login(user User_db) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getuser := User_db{}
	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&getuser)
	if err != nil {
		return nil, fmt.Errorf("Not found username :'%s'", user.Username)
	}
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

	objId, _ := primitive.ObjectIDFromHex(user_id)
	err := collection.FindOne(ctx, bson.M{"userID": objId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *userRepositoryDB) Update(user_id string, updatedUser User_db) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user_db, _ := r.GetById(user_id)
	version := user_db.Version
	for updatedUser.Version != version {
		user_db, _ = r.GetById(user_id)
		updatedUser.Version = user_db.Version
	}

	updatedUser.Version = updatedUser.Version + 1
	update := bson.M{"firstname": updatedUser.Firstname, "lastname": updatedUser.Lastname,
		"birthdate": updatedUser.Birthdate, "programing_skill": updatedUser.Programing_skill,
		"version": updatedUser.Version}
	objId, _ := primitive.ObjectIDFromHex(user_id)
	result, err := collection.UpdateOne(ctx, bson.M{"userID": objId}, bson.M{"$set": update})
	if err != nil {
		return &updatedUser, err
	}
	if result.MatchedCount == 1 {
		err := collection.FindOne(ctx, bson.M{"userID": objId}).Decode(&updatedUser)
		if err != nil {
			return &updatedUser, err
		}
		return &updatedUser, err
	}
	return &updatedUser, err

}

func (r *userRepositoryDB) Delete(user_id string) (*User_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(user_id)
	user_db, _ := r.GetById(user_id)

	// filter := bson.M{"$and": []interface{}{bson.M{"_id": objId}, bson.M{"username": username}}}
	result, err := collection.DeleteOne(ctx, bson.M{"userID": objId})
	if err != nil {
		return user_db, err
	}
	if result.DeletedCount < 1 {
		return user_db, fmt.Errorf("Not found User_id :'%s'", user_id)
	}
	return user_db, err

}
