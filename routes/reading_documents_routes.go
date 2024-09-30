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
	mangaURLPath := r.Group(commonPath + "/reading-docs/mangas")

	mangaURLPath.GET("/:id", bookController.GetMangaById())
	mangaURLPath.POST("/", bookController.CreateManga())
	mangaURLPath.PUT("/:id", bookController.UpdateManga())
	mangaURLPath.DELETE("/:id", bookController.DeleteManga())
}

func DocumentRoutes(r *gin.Engine, documentController controllers.DocumentController) {
	documentURLPath := r.Group(commonPath + "/reading-docs/documents")

	documentURLPath.GET("/:id", documentController.GetDocumentById())
	documentURLPath.POST("/", documentController.CreateDocument())
	documentURLPath.PUT("/:id", documentController.UpdateDocument())
	documentURLPath.DELETE("/:id", documentController.DeleteDocument())
}

func ArticleRoutes(r *gin.Engine, articleController controllers.ArticleController) {
	articleURLPath := r.Group(commonPath + "/reading-docs/articles")

	articleURLPath.GET("/:id", articleController.GetArticleById())
	articleURLPath.POST("/", articleController.CreateArticle())
	articleURLPath.PUT("/:id", articleController.UpdateArticle())
	articleURLPath.DELETE("/:id", articleController.DeleteArticle())
}
