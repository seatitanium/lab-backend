package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Respond(c *gin.Context, ok bool, note string, message string, data any) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok":      ok,
		"note":    note,
		"message": message,
		"data":    data,
	})
}

func RespondOk(c *gin.Context, note string, message string, data any) {
	Respond(c, true, note, message, nil)
}

func RespondNg(c *gin.Context, note string, message string, data any) {
	Respond(c, false, note, message, nil)
}
