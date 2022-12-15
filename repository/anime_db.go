package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type animeRepositoryDB struct {
	db *mongo.Client
}

func NewAnimeRepositoryDB(db *mongo.Client) AnimeRepository {
	return &animeRepositoryDB{db: db}
}

func (r *animeRepositoryDB) getCollection() *mongo.Collection {
	collection := r.db.Database(DBName_anime).Collection(collection_anime)
	return collection
}

func (r *animeRepositoryDB) Create(anime Anime_db) (*Anime_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// // check title
	filter := bson.D{
		{Key: "$or",
			Value: bson.A{
				bson.M{"title": anime.Title},
				bson.M{"title_th": anime.Title_th},
			},
		},
	}
	err := collection.FindOne(ctx, filter).Decode(&anime)
	// if found title
	if err == nil {
		return nil, fmt.Errorf("title or title_th not ready")
	}

	// create and check animeid
	anime_id := "animeID" + uuid.New().String()
	err = collection.FindOne(ctx, bson.M{"anime_ID": anime_id}).Decode(&anime)
	anime.Anime_ID = anime_id
	for err == nil {
		anime_id := "anime_id" + uuid.New().String()
		err = collection.FindOne(ctx, bson.M{"anime_ID": anime_id}).Decode(&anime)
		anime.Anime_ID = anime_id
	}

	_, err = collection.InsertOne(ctx, &anime)

	return &anime, err

}

func (r *animeRepositoryDB) Getall() ([]Anime_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	animes := []Anime_db{}
	results, err := collection.Find(ctx, bson.M{})
	defer results.Close(ctx)
	for results.Next(ctx) {
		singleAnime := Anime_db{}
		err = results.Decode(&singleAnime)
		animes = append(animes, singleAnime)
	}
	if err != nil {
		return nil, err
	}
	return animes, err
}

func (r *animeRepositoryDB) Delete(anime_id string) (*Anime_db, error) {
	collection := r.getCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check animeID
	getAnime := Anime_db{}
	filter := bson.M{"anime_ID": anime_id}
	err := collection.FindOne(ctx, filter).Decode(&getAnime)
	if err != nil {
		return nil, err
	}
	result, err := collection.DeleteOne(ctx, bson.M{"anime_ID": anime_id})
	if err != nil {
		return nil, err
	}
	if result.DeletedCount < 1 {
		return nil, fmt.Errorf("Not found anime_id")
	}
	return &getAnime, err
}
