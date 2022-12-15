package service

import "go_test/repository"

type AnimeService interface {
	CreateAnime(newAnime AnimeRequest) (*repository.Anime_db, error)
	GetAllAnime() ([]repository.Anime_db, error)
	DeleteAnime(anime_id string) (*repository.Anime_db, error)
}

type AnimeRequest struct {
	Src      string `json:"src" bson:"src" validate:"required"`
	Title    string `json:"title" bson:"title"  validate:"required"`
	Title_th string `json:"title_th" bson:"title_th"  validate:"required"`
	Trailer  string `json:"trailer" bson:"trailer"  validate:"required"`
}
