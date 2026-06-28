package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/internal/controllers"
	"github.com/gin-gonic/gin"
)

const commonPath = "/v1/api"

func UserRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, userController controllers.UserController) {
	// :id = readingId
	usersPath := r.Group(commonPath + "/users")
	usersPath.GET("/:id", rateLimiter, userController.GetUserById())
	usersPath.POST("/", rateLimiter, userController.CreateUser())
	usersPath.PUT("/:id", rateLimiter, userController.UpdateUser())
	usersPath.DELETE("/:id", rateLimiter, userController.DeleteUser())
}
