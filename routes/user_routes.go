package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

const commonPath = "/v1/api"

func UserRoutes(r *gin.Engine, userController controllers.UserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/users")
	usersPath.GET("/:id", userController.GetUserById())
	usersPath.POST("/", userController.CreateUser())
	usersPath.PUT("/:id", userController.UpdateUser())
	usersPath.DELETE("/:id", userController.DeleteUser())
}

func ReadingUserRoutes(r *gin.Engine, userReadingController controllers.ReadingUserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/user-readings")
	usersPath.GET("/", userReadingController.GetMyReadings())
	usersPath.POST("/", userReadingController.StartReading())
	usersPath.PUT("/:id", userReadingController.UpdateMyReading())
	usersPath.DELETE("/:id", userReadingController.DeleteMyReading())
}

func RecordUserRoutes(r *gin.Engine, userReadingController controllers.RecordUserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/user-records")
	usersPath.GET("/:id", userReadingController.GetRecordsFromMyReading())
	usersPath.POST("/", userReadingController.AddRecord())
	usersPath.PUT("/:id", userReadingController.UpdateRecord())
	usersPath.DELETE("/:id", userReadingController.RemoveMyRecord())
}
