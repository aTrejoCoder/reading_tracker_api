package main

import (
	"github.com/aTrejoCoder/reading_tracker_api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	ratelimiterHandler := middleware.RateLimiter()

	// Server
	r := gin.Default()
	r.GET("/home", ratelimiterHandler, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"home": "reading_tracker_api"})
	})

	r.Run()
}
