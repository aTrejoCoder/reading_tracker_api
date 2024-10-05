package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine, bookController controllers.BookController) {
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

func MangaRoutes(r *gin.Engine, mangaController controllers.MangaController) {
	mangaURLPath := r.Group(commonPath + "/reading-docs/mangas")

	mangaURLPath.GET("/:id", mangaController.GetMangaById())
	mangaURLPath.GET("by-name/:name", mangaController.GetMangaByMatchingName())
	mangaURLPath.GET("by-author/:author", mangaController.GetMangaByAuthor())
	mangaURLPath.GET("by-genre/:genre", mangaController.GetMangaByGenre())
	mangaURLPath.GET("by-demography/:demography", mangaController.GetMangaByDemography())
	mangaURLPath.GET("/all", mangaController.GetAllMangasSortedPaginated())

	mangaURLPath.POST("/", mangaController.CreateManga())
	mangaURLPath.PUT("/:id", mangaController.UpdateManga())
	mangaURLPath.DELETE("/:id", mangaController.DeleteManga())
}

func CustomDocumentUserRoutes(r *gin.Engine, documentController controllers.DocumentController) {
	documentURLPath := r.Group(commonPath + "/user/reading-docs/custom-documents")

	documentURLPath.GET("/my-docs", documentController.GetMyCustomDocuments())
	documentURLPath.GET("/:id", documentController.GetDocumentById())
	documentURLPath.POST("/", documentController.CreateDocument())
	documentURLPath.PUT("/:id", documentController.UpdateDocument())
	documentURLPath.DELETE("/:id", documentController.DeleteDocument())
}
