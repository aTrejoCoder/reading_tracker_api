package routes

import (
	"github.com/aTrejoCoder/reading_tracker_api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authController controllers.AuthController) {
	authRoutes := r.Group(commonPath)
	authRoutes.POST("/signup", authController.Signup())
	authRoutes.POST("/login", authController.Login())
}
