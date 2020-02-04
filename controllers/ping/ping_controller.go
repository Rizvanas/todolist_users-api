package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping checks if connection is live and returns pong!
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong!")
}
