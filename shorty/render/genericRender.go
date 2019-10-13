package render

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render one of JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, JSON is returned
func Render(c *gin.Context, data gin.H) {

	errorPublic := c.Errors.ByType(gin.ErrorTypePublic)
	errorPrivate := c.Errors.ByType(gin.ErrorTypePublic)

	if errorPublic != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": errorPublic,
		})
	} else if errorPrivate != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "something went wrong, we are working on it",
		})
	}

	switch c.Request.Header.Get("Accept") {

	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	}
}
