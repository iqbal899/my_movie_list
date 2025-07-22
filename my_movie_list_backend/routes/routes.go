package routes

import (
	"github.com/gofiber/fiber/v2"
	"my_movie_list_backend/controllers"
	"my_movie_list_backend/middleware"
)

func SetupRoutes(app *fiber.App) {

	// User Registration and Login
	app.Post("/register", controllers.RegisterUser)  // Register a new user
	app.Post("/login", controllers.LoginUser)        // Login and get JWT token

	// Movie Search (uses OMDb API, no login required)
	app.Get("/search", controllers.SearchMovie)     // Search for movies using query param


	// Group all authenticated routes under "/api"
	api := app.Group("/api", middleware.AuthMiddleware)
	api.Post("/watchlist",  controllers.AddToWatchList)
api.Get("/watchlist",  controllers.GetWatchList)
api.Patch("/watchlist/:id/watched", controllers.MarkAsWatched)
api.Post("/watchlist/:id/review",  controllers.AddReview)
api.Delete("/watchlist/:id", controllers.DeleteFromWatchList)


}
