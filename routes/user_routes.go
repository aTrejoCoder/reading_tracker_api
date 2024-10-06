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
