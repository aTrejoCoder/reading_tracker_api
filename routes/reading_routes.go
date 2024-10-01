package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func ReadingRoutes(r *gin.Engine, bookController controllers.ReadingController) {
	mangaURLPath := r.Group(commonPath + "/readings")

	mangaURLPath.GET("/:id", bookController.GetReadingById())
	mangaURLPath.POST("/", bookController.CreateReading())
	mangaURLPath.PUT("/:id", bookController.UpdateReading())
	mangaURLPath.DELETE("/:id", bookController.DeleteReading())
}
