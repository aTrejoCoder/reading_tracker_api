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

func RecordRoutes(r *gin.Engine, record controllers.RecordController) {
	readingRecordPath := r.Group(commonPath + "/readings/records")

	readingRecordPath.POST("/", record.CreateRecord())
	readingRecordPath.PUT("/:id", record.UpdateRecord())
	readingRecordPath.DELETE("/:id", record.DeleteRecord())
}

func ReadinListRoutes(r *gin.Engine, list controllers.ReadingListController) {
	listPath := r.Group(commonPath + "/readings/lists")

	listPath.PUT("/add-readings/:id", list.AddReadingToList())
	listPath.PUT("/remove-readings/:id", list.RemoveReadingToList())

	listPath.GET("/", list.GetReadingListByUserId())
	listPath.POST("/", list.CreateReadingList())
	listPath.PUT("/:id", list.UpdateReadingList())
	listPath.DELETE("/:id", list.DeleteReadingList())

}
