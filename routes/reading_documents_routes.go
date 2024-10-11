package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, bookController controllers.BookController) {
	bookURLPath := r.Group(commonPath + "/reading-docs/books")

	bookURLPath.GET("/:id", bookController.GetBookById())
	bookURLPath.GET("/by-name/:name", bookController.GetBooksByMatchingName())
	bookURLPath.GET("/by-isbn/:isbn", bookController.GetBookByISBN())
	bookURLPath.GET("/by-author/:author", bookController.GetBooksByAuthor())
	bookURLPath.GET("/by-genre/:genre", bookController.GetBooksByGenre())
	bookURLPath.GET("/all", bookController.GetAllBooksSortedPaginated())

	bookURLPath.POST("/", bookController.CreateBook())
	bookURLPath.PUT("/:id", bookController.UpdateBook())
	bookURLPath.DELETE("/:id", bookController.DeleteBook())
}

func MangaRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, mangaController controllers.MangaController) {
	mangaURLPath := r.Group(commonPath + "/reading-docs/mangas")

	mangaURLPath.GET("/:id", rateLimiter, mangaController.GetMangaById())
	mangaURLPath.GET("by-name/:name", rateLimiter, mangaController.GetMangaByMatchingName())
	mangaURLPath.GET("by-author/:author", rateLimiter, mangaController.GetMangaByAuthor())
	mangaURLPath.GET("by-genre/:genre", rateLimiter, mangaController.GetMangaByGenre())
	mangaURLPath.GET("by-demography/:demography", rateLimiter, mangaController.GetMangaByDemography())
	mangaURLPath.GET("/all", rateLimiter, mangaController.GetAllMangasSortedPaginated())

	mangaURLPath.POST("/", rateLimiter, mangaController.CreateManga())
	mangaURLPath.PUT("/:id", rateLimiter, mangaController.UpdateManga())
	mangaURLPath.DELETE("/:id", rateLimiter, mangaController.DeleteManga())
}

func CustomDocumentUserRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, documentController controllers.DocumentController) {
	documentURLPath := r.Group(commonPath + "/user/reading-docs/custom-documents")

	documentURLPath.GET("/my-docs", rateLimiter, documentController.GetMyCustomDocuments())
	documentURLPath.GET("/:id", rateLimiter, documentController.GetDocumentById())
	documentURLPath.POST("/", rateLimiter, documentController.CreateDocument())
	documentURLPath.PUT("/:id", rateLimiter, documentController.UpdateDocument())
	documentURLPath.DELETE("/:id", rateLimiter, documentController.DeleteDocument())
}
