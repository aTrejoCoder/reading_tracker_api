package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, rateLimiter gin.HandlerFunc, authController controllers.AuthController) {
	authRoutes := r.Group(commonPath)
	authRoutes.POST("/signup", rateLimiter, authController.Signup())
	authRoutes.POST("/login", rateLimiter, authController.Login())
}
