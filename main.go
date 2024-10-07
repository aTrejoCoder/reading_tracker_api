package main

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/aTrejoCoder/reading_tracker_api/database"
	_ "github.com/aTrejoCoder/reading_tracker_api/docs"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/routes"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Reading Tracker API
// @version 1.0
// @description API for managing books, manga, readings, and custom documents.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

func main() {
	// Server
	r := gin.Default()
	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"home": "reading_tracker_api"})
	})

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Database
	database.DbConn()
	userCollection := database.Client.Database("reading_tracker").Collection("users")
	bookCollection := database.Client.Database("reading_tracker").Collection("books")
	mangaCollection := database.Client.Database("reading_tracker").Collection("mangas")
	readingCollection := database.Client.Database("reading_tracker").Collection("readings")

	// Repository
	commonUserRepository := repository.NewRepository[models.User](userCollection)
	commonBookRepository := repository.NewRepository[models.Book](bookCollection)
	commonMangaRepository := repository.NewRepository[models.Manga](mangaCollection)
	commonDocumentRepository := repository.NewCustomDocumentRepository(*userCollection, *commonUserRepository)

	commonReadingRepository := repository.NewRepository[models.Reading](readingCollection)

	readingListRepository := repository.NewReadingListRepository(*userCollection, *commonUserRepository)
	readingExtendService := repository.NewReadingExtendRepository(*readingCollection)
	userRepository := repository.NewUserRepository(userCollection)

	// Service
	userService := services.NewUserService(*commonUserRepository, userRepository)
	authService := services.NewAuthService(userRepository, *commonUserRepository)
	bookService := services.NewBookService(*commonBookRepository)
	mangaService := services.NewMangaService(*commonMangaRepository)

	readingListService := services.NewReadingListService(*readingListRepository)
	readingService := services.NewReadingService(*commonReadingRepository, *readingExtendService, *commonDocumentRepository, *commonMangaRepository, *commonBookRepository, *commonUserRepository)
	readingRecordService := services.NewReadingRecordService(*commonReadingRepository, *readingExtendService)
	documentService := services.NewCustomDocumentService(*commonDocumentRepository)

	// Controller
	userControler := controllers.NewUserController(userService)

	authController := controllers.NewAuthController(authService)
	bookController := controllers.NewBookController(bookService)
	mangaController := controllers.NewMangaController(mangaService)

	readingListController := controllers.NewReadingListController(readingListService)
	readingListUserController := controllers.NewReadingListUserController(readingListService)
	readingController := controllers.NewReadingControler(readingService)
	readingRecordController := controllers.NewRecordController(readingRecordService)
	readingUserController := controllers.NewReadingUserController(readingService)

	documentController := controllers.NewDocumentController(documentService)
	recordUserController := controllers.NewRecordUserController(readingRecordService)

	// Routes
	routes.UserRoutes(r, *userControler)
	routes.AuthRoutes(r, *authController)
	routes.BookRoutes(r, *bookController)
	routes.MangaRoutes(r, *mangaController)

	routes.ReadingRoutes(r, *readingController)
	routes.RecordRoutes(r, *readingRecordController)
	routes.ReadingUserRoutes(r, *readingUserController)
	routes.RecordUserRoutes(r, *recordUserController)
	routes.ReadingListRoutes(r, *readingListController)
	routes.ReadingListUserRoutes(r, *readingListUserController)
	routes.CustomDocumentUserRoutes(r, *documentController)

	r.Run()
}
