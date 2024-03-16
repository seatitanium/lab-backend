package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 向前端返回 JSON 格式的信息
//   - c - *gin.Context
//   - ok - 表示操作是否成功完成
//   - note - 指示信息，用于精确描述操作执行结果
//   - message - 面向用户的提示信息
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
