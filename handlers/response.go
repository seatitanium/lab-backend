package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seatimc/backend/errors"
)

type Response struct {
	HttpCode int         `json:"-"`
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
}

func RespSuccess(ctx *gin.Context, data ...interface{}) {
	resp := Response{
		HttpCode: http.StatusOK,
		Code:     0,
		Msg:      "Success",
		Data:     nil,
	}

	switch len(data) {
	case 1:
		resp.Data = data[0]
	default:
		resp.Data = nil
	}
	ctx.JSON(resp.HttpCode, resp)
}

func RespTokenError(ctx *gin.Context, errCode int, errMsg string) {
	resp := Response{
		HttpCode: http.StatusUnauthorized,
		Code:     errCode,
		Msg:      errMsg,
		Data:     false,
	}
	ctx.JSON(resp.HttpCode, resp)
}

func RespForbidden(ctx *gin.Context) {
	resp := Response{
		HttpCode: http.StatusForbidden,
		Code:     errors.RespErrCodeForbidden,
		Msg:      errors.RespErrMsgForbidden,
		Data:     false,
	}
	ctx.JSON(resp.HttpCode, resp)
}
