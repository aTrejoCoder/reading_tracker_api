package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Server
	r := gin.Default()
	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"home": "reading_tracker_api"})
	})

	r.Run()
}
