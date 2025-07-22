package models

type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Genre      string `json:"Genre"`
	Poster     string `json:"Poster"`
	IMDBRating string `json:"imdbRating"`
}
