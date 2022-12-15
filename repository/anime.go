package repository

type AnimeRepository interface {
	Create(anime Anime_db) (*Anime_db, error)
	Getall() ([]Anime_db, error)
	Delete(anime_id string) (*Anime_db, error)
}

const (
	DBName_anime     = "kong_test" //kong_test
	collection_anime = "animes"    //user2
)

type Anime_db struct {
	Anime_ID string `json:"anime_ID" bson:"anime_ID"`
	Src      string `json:"src" bson:"src"`
	Title    string `json:"title" bson:"title"`
	Title_th string `json:"title_th" bson:"title_th"`
	Trailer  string `json:"trailer" bson:"trailer"`
}
