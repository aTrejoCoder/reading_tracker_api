package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

// Readings
func ReadingRoutes(r *gin.Engine, reading controllers.ReadingController) {
	readingPath := r.Group(commonPath + "/readings")

	readingPath.GET("/:id", reading.GetReadingById())
	readingPath.POST("/", reading.CreateReading())
	readingPath.PUT("/:id", reading.UpdateReading())
	readingPath.DELETE("/:id", reading.DeleteReading())
}

func ReadingUserRoutes(r *gin.Engine, userReadingController controllers.ReadingUserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/user-readings")
	usersPath.GET("/", userReadingController.GetMyReadings())
	usersPath.POST("/", userReadingController.StartReading())
	usersPath.PUT("/:id", userReadingController.UpdateMyReading())
	usersPath.DELETE("/:id", userReadingController.DeleteMyReading())
}

// Reading Records
func RecordRoutes(r *gin.Engine, record controllers.RecordController) {
	readingRecordPath := r.Group(commonPath + "/readings/records")

	readingRecordPath.POST("/", record.CreateRecord())
	readingRecordPath.PUT("/:id", record.UpdateRecord())
	readingRecordPath.DELETE("/:id", record.DeleteRecord())
}

func RecordUserRoutes(r *gin.Engine, userReadingController controllers.RecordUserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/user-records")
	usersPath.GET("/:id", userReadingController.GetRecordsFromMyReading())
	usersPath.POST("/", userReadingController.AddRecord())
	usersPath.PUT("/:id", userReadingController.UpdateRecord())
	usersPath.DELETE("/:id", userReadingController.RemoveMyRecord())
}

// Reading List
func ReadingListRoutes(r *gin.Engine, list controllers.ReadingListController) {
	listPath := r.Group(commonPath + "/readings/lists")

	listPath.PUT("/add-readings/:id", list.AddReadingToList())
	listPath.PUT("/remove-readings/:id", list.RemoveReadingToList())

	listPath.GET("/", list.GetReadingListByUserId())
	listPath.POST("/", list.CreateReadingList())
	listPath.PUT("/:id", list.UpdateReadingList())
	listPath.DELETE("/:id", list.DeleteReadingList())

}

func ReadingListUserRoutes(r *gin.Engine, list controllers.ReadingListUserController) {
	listPath := r.Group(commonPath + "/user-readings/lists")

	listPath.PUT("/add-readings", list.AddReadingToList())
	listPath.PUT("/remove-readings/", list.RemoveReadingToList())

	listPath.GET("/:id", list.GetMyReadingListById())
	listPath.GET("/all", list.GetMyReadingLists())
	listPath.POST("/", list.CreateReadingList())
	listPath.PUT("/:id", list.UpdateMyReadingList())
	listPath.DELETE("/:id", list.DeleteMyReadingList())

}
