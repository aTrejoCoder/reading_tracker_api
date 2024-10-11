package middleware

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(429, gin.H{
		"error":        "Too many requests",
		"message":      "You have exceeded the request limit.",
		"retry_after":  time.Until(info.ResetTime).Seconds(),
		"limit":        info.Limit,
		"reset_time":   info.ResetTime.Format(time.RFC3339),
		"current_time": time.Now().Format(time.RFC3339),
	})
}

func RateLimiter() gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 30,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	return mw
}
