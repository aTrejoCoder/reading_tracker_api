package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Readings
func ReadingRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, reading controllers.ReadingController) {
	readingPath := r.Group(commonPath + "/readings")

	readingPath.GET("/:id", reading.GetReadingById())
	readingPath.POST("/", reading.CreateReading())
	readingPath.PUT("/:id", reading.UpdateReading())
	readingPath.DELETE("/:id", reading.DeleteReading())
}

func ReadingUserRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, userReadingController controllers.ReadingUserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/user-readings")
	usersPath.GET("/", rateLimiter, userReadingController.GetMyReadings())
	usersPath.GET("/by-status", rateLimiter, userReadingController.GetMyReadingsByStatus())
	usersPath.GET("/by-type", rateLimiter, userReadingController.GetMyReadingsByType())
	usersPath.POST("/", rateLimiter, userReadingController.StartReading())
	usersPath.PUT("/:id", rateLimiter, userReadingController.UpdateMyReading())
	usersPath.DELETE("/:id", rateLimiter, userReadingController.DeleteMyReading())
}

// Reading Records
func RecordRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, record controllers.RecordController) {
	readingRecordPath := r.Group(commonPath + "/readings/records")

	readingRecordPath.POST("/", rateLimiter, record.CreateRecord())
	readingRecordPath.PUT("/:id", rateLimiter, record.UpdateRecord())
	readingRecordPath.DELETE("/:id", rateLimiter, record.DeleteRecord())
}

func RecordUserRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, userReadingController controllers.RecordUserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/user-records")
	usersPath.GET("/:id", rateLimiter, userReadingController.GetRecordsFromMyReading())
	usersPath.POST("/", rateLimiter, userReadingController.AddRecord())
	usersPath.PUT("/:id", rateLimiter, userReadingController.UpdateRecord())
	usersPath.DELETE("/:id", rateLimiter, userReadingController.RemoveMyRecord())
}

// Reading List
func ReadingListRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, list controllers.ReadingListController) {
	listPath := r.Group(commonPath + "/readings/lists")

	listPath.PUT("/add-readings/:id", rateLimiter, list.AddReadingToList())
	listPath.PUT("/remove-readings/:id", rateLimiter, list.RemoveReadingToList())

	listPath.GET("/", rateLimiter, list.GetReadingListByUserId())
	listPath.POST("/", rateLimiter, list.CreateReadingList())
	listPath.PUT("/:id", rateLimiter, list.UpdateReadingList())
	listPath.DELETE("/:id", rateLimiter, list.DeleteReadingList())

}

func ReadingListUserRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, list controllers.ReadingListUserController) {
	listPath := r.Group(commonPath + "/user-readings/lists")

	listPath.PUT("/add-readings", rateLimiter, list.AddReadingToList())
	listPath.PUT("/remove-readings/", rateLimiter, list.RemoveReadingToList())

	listPath.GET("/:id", rateLimiter, list.GetMyReadingListById())
	listPath.GET("/all", rateLimiter, list.GetMyReadingLists())
	listPath.POST("/", rateLimiter, list.CreateReadingList())
	listPath.PUT("/:id", rateLimiter, list.UpdateMyReadingList())
	listPath.DELETE("/:id", rateLimiter, list.DeleteMyReadingList())

}
