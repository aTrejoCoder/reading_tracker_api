package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine, bookController controllers.BookController) {
	bookURLPath := r.Group(commonPath + "/reading-docs/books")

	bookURLPath.GET("/:id", bookController.GetBookById())
	bookURLPath.POST("/", bookController.CreateBook())
	bookURLPath.PUT("/:id", bookController.UpdateBook())
	bookURLPath.DELETE("/:id", bookController.DeleteBook())
}

func MangaRoutes(r *gin.Engine, bookController controllers.MangaController) {
	bookURLPath := r.Group(commonPath + "/reading-docs/mangas")

	bookURLPath.GET("/:id", bookController.GetMangaById())
	bookURLPath.POST("/", bookController.CreateManga())
	bookURLPath.PUT("/:id", bookController.UpdateManga())
	bookURLPath.DELETE("/:id", bookController.DeleteManga())
}

func DocumentRoutes(r *gin.Engine, documentController controllers.DocumentController) {
	bookURLPath := r.Group(commonPath + "/reading-docs/documents")

	bookURLPath.GET("/:id", documentController.GetDocumentById())
	bookURLPath.POST("/", documentController.CreateDocument())
	bookURLPath.PUT("/:id", documentController.UpdateDocument())
	bookURLPath.DELETE("/:id", documentController.DeleteDocument())
}
