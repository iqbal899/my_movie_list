package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"my_movie_list_backend/database"
	"my_movie_list_backend/models"
	"net/http"
	"os"
)

func AddToWatchList(c *fiber.Ctx) error {
	var body struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userID := c.Locals("user_id").(uint)
	apiKey := os.Getenv("OMDB_API_KEY")

	res, err := req.C().R().SetQueryParams(map[string]string{
		"t":      body.Title,
		"apikey": apiKey,
	}).SetSuccessResult(&models.Movie{}).Get("https://www.omdbapi.com/")

	if err != nil || !res.IsSuccess() {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"error": "OMDb fetch failed"})
	}
	movie := res.Result().(*models.Movie)

	entry := models.WatchList{
		ID:         uuid.New().String(),
		UserID:     fmt.Sprint(userID),
		Title:      movie.Title,
		Year:       movie.Year,
		Genre:      movie.Genre,
		Poster:     movie.Poster,
		IMDBRating: movie.IMDBRating,
		Watched:    false,
	}
	if err := database.DB.Create(&entry).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "DB Error"})
	}
	return c.Status(http.StatusCreated).JSON(entry)
}

func GetWatchList(c *fiber.Ctx) error {
	userID := fmt.Sprint(c.Locals("user_id").(uint))
	var list []models.WatchList
	if err := database.DB.Where("user_id = ?", userID).Find(&list).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "DB Error"})
	}
	return c.JSON(list)
}

func MarkAsWatched(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := fmt.Sprint(c.Locals("user_id").(uint))

	var entry models.WatchList
	if err := database.DB.First(&entry, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}
	entry.Watched = true
	database.DB.Save(&entry)
	return c.JSON(fiber.Map{"message": "Marked as watched"})
}

func AddReview(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := fmt.Sprint(c.Locals("user_id").(uint))

	var body struct {
		Review string  `json:"review"`
		Rating float64 `json:"rating"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var entry models.WatchList
	if err := database.DB.First(&entry, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}
	entry.Review = body.Review
	entry.Rating = body.Rating
	database.DB.Save(&entry)
	return c.JSON(fiber.Map{"message": "Review added"})
}

func DeleteFromWatchList(c *fiber.Ctx) error {
	id := c.Params("id")
	userID := fmt.Sprint(c.Locals("user_id").(uint))

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.WatchList{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete failed"})
	}
	return c.JSON(fiber.Map{"message": "Deleted"})
}
