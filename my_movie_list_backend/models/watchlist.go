package models

type WatchList struct {
	ID         string  `gorm:"primaryKey" json:"id"`
	UserID     string  `json:"user_id"`
	Title      string  `json:"title"`
	Year       string  `json:"year"`
	Genre      string  `json:"genre"`
	Poster     string  `json:"poster"`
	IMDBRating string  `json:"imdb_rating"`
	Watched    bool    `json:"watched"`
	Review     string  `json:"review"`
	Rating     float64 `json:"rating"`
}
