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
	mangaCollection := database.Client.Database("reading_tracker").Collection("mangas")
	documentCollection := database.Client.Database("reading_tracker").Collection("documents")
	articleCollection := database.Client.Database("reading_tracker").Collection("article")
	readingCollection := database.Client.Database("reading_tracker").Collection("readings")

	// Repository
	commonUserRepository := repository.NewRepository[models.User](userCollection)
	commonBookRepository := repository.NewRepository[models.Book](bookCollection)
	commonMangaRepository := repository.NewRepository[models.Manga](mangaCollection)
	commonDocumentRepository := repository.NewRepository[models.Document](documentCollection)
	commonArticleRepository := repository.NewRepository[models.Article](articleCollection)
	commonReadingRepository := repository.NewRepository[models.Reading](readingCollection)

	readingListRepository := repository.NewReadingListRepository(*userCollection, *commonUserRepository)
	readingExtendService := repository.NewReadingExtendRepository(*readingCollection)
	userRepository := repository.NewUserRepository(userCollection)

	// Service
	userService := services.NewUserService(*commonUserRepository, userRepository)
	authService := services.NewAuthService(userRepository, *commonUserRepository)
	bookService := services.NewBookService(*commonBookRepository)
	mangaService := services.NewMangaService(*commonMangaRepository)
	documentService := services.NewDocumentService(*commonDocumentRepository)
	articleService := services.NewArticleService(*commonArticleRepository)

	readingListService := services.NewReadingListService(*readingListRepository)
	readingService := services.NewReadingService(*commonReadingRepository, *readingExtendService, userRepository)
	readingRecordService := services.NewReadingRecordService(*commonReadingRepository, *readingExtendService)

	// Controller
	userControler := controllers.NewUserController(userService)

	authController := controllers.NewAuthController(authService)
	bookController := controllers.NewBookController(bookService)
	mangaController := controllers.NewMangaController(mangaService)
	documentController := controllers.NewDocumentController(documentService)
	articleController := controllers.NewArticleController(articleService)

	readingListController := controllers.NewReadingListController(readingListService)
	readingListUserController := controllers.NewReadingListUserController(readingListService)
	readingController := controllers.NewReadingControler(readingService)
	readingRecordController := controllers.NewRecordController(readingRecordService)
	readingUserController := controllers.NewReadingUserController(readingService)

	recordUserController := controllers.NewRecordUserController(readingRecordService)

	// Routes
	routes.UserRoutes(r, *userControler)
	routes.AuthRoutes(r, *authController)
	routes.BookRoutes(r, *bookController)
	routes.MangaRoutes(r, *mangaController)
	routes.DocumentRoutes(r, *documentController)
	routes.ArticleRoutes(r, *articleController)

	routes.ReadingRoutes(r, *readingController)
	routes.RecordRoutes(r, *readingRecordController)
	routes.ReadingUserRoutes(r, *readingUserController)
	routes.RecordUserRoutes(r, *recordUserController)
	routes.ReadingListRoutes(r, *readingListController)
	routes.ReadingListUserRoutes(r, *readingListUserController)

	r.Run()
}
