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
	bookCollection := database.Client.Database("reading_tracker").Collection("bokks")

	// Repository
	commonUserRepository := repository.NewRepository[models.User](userCollection)
	commonBookRepository := repository.NewRepository[models.Book](bookCollection)
	userRepository := repository.NewUserRepository(userCollection)

	// Service
	userService := services.NewUserService(*commonUserRepository, userRepository)
	authService := services.NewAuthService(userRepository, *commonUserRepository)
	bookService := services.NewBookService(*commonBookRepository)

	// Controller
	userControler := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	bookController := controllers.NewBookController(bookService)

	// Routes
	routes.UserRoutes(r, *userControler)
	routes.AuthRoutes(r, *authController)
	routes.BookRoutes(r, *bookController)

	r.Run()
}
