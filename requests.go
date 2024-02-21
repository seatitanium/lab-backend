package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Respond(c *gin.Context, ok bool, message string, data any) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok":      ok,
		"message": message,
		"data":    data,
	})
}
