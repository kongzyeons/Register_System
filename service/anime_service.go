package service

import (
	"go_test/repository"
	"log"
)

type animeService struct {
	animeRepo repository.AnimeRepository
}

func NewAnimeService(animeRepo repository.AnimeRepository) AnimeService {
	return &animeService{animeRepo: animeRepo}
}

func (s *animeService) CreateAnime(newAnime AnimeRequest) (*repository.Anime_db, error) {
	animeCreate := repository.Anime_db{
		Src:      newAnime.Src,
		Title:    newAnime.Title,
		Title_th: newAnime.Title_th,
		Trailer:  newAnime.Trailer,
	}

	user_db, err := s.animeRepo.Create(animeCreate)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user_db, err
}

func (s *animeService) GetAllAnime() ([]repository.Anime_db, error) {
	animes, err := s.animeRepo.Getall()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return animes, err
}

func (s *animeService) DeleteAnime(anime_id string) (*repository.Anime_db, error) {
	anime, err := s.animeRepo.Delete(anime_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return anime, err
}
