package controllers

import (
	"encoding/json"
	"fmt"
	"my_movie_list_backend/models"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SearchMovie(c *fiber.Ctx) error {
	title := c.Query("title")
	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title query is required",
		})
	}

	apiKey := os.Getenv("OMDB_API_KEY")
	url := fmt.Sprintf("https://www.omdbapi.com/?t=%s&apikey=%s", title, apiKey)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": "Failed to fetch from OMDb",
		})
	}
	defer resp.Body.Close()

	var movie models.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil || movie.Title == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Movie not found",
		})
	}

	return c.JSON(movie)
}
