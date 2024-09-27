package main

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/aTrejoCoder/reading_tracker_api/database"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/routes"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Server
	r := gin.Default()
	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"home": "reading_tracker_api"})
	})

	//Database
	database.DbConn()
	userCollection := database.Client.Database("reading_tracker").Collection("users")

	// Repository
	commonUserRepository := repository.NewRepository[models.User](userCollection)
	userRepository := repository.NewUserRepository(userCollection)

	// Service
	userService := services.NewUserService(*commonUserRepository, userRepository)

	// Controller
	userControler := controllers.NewUserController(userService)

	// Routes
	routes.UserRoutes(r, *userControler)

	r.Run()
}
