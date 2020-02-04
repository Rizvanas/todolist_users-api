package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication maps urls to handling functions and stars server on port :8080
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
