package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func ReadingRoutes(r *gin.Engine, reading controllers.ReadingController) {
	readingPath := r.Group(commonPath + "/readings")

	readingPath.GET("/:id", reading.GetReadingById())
	readingPath.POST("/", reading.CreateReading())
	readingPath.PUT("/:id", reading.UpdateReading())
	readingPath.DELETE("/:id", reading.DeleteReading())
}

func RecordRoutes(r *gin.Engine, record controllers.ReadingRecordController) {
	readingRecordPath := r.Group(commonPath + "/readings/records")

	readingRecordPath.POST("/", record.CreateRecord())
	readingRecordPath.PUT("/:id", record.UpdateRecord())
	readingRecordPath.DELETE("/:id", record.DeleteRecord())
}
