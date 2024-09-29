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
