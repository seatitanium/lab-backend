package bss

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"strconv"
	"time"
)

func HandleQueryBill(ctx *gin.Context) *errHandler.CustomErr {
	pagenum := ctx.Query("pagenum")
	pagesize := ctx.Query("pagesize")

	var _pagenumN int64
	var _pagesizeN int64
	var err error

	if pagenum == "" {
		_pagenumN = 1
	} else {
		_pagenumN, err = strconv.ParseInt(pagenum, 10, 32)

		if err != nil {
			return errHandler.WrongParam()
		}

		if _pagenumN <= 0 {
			return errHandler.WrongParam()
		}
	}

	if pagesize == "" {
		_pagesizeN = 10
	} else {
		_pagesizeN, err = strconv.ParseInt(pagesize, 10, 32)

		if err != nil {
			return errHandler.WrongParam()
		}

		if _pagesizeN <= 0 {
			return errHandler.WrongParam()
		}
	}

	var pagenumN = int32(_pagenumN)
	var pagesizeN = int32(_pagesizeN)

	result, customErr := bss.QueryBill(fmt.Sprintf("%d-%02d", time.Now().Year(), time.Now().Month()), pagenumN, pagesizeN)

	if err != nil {
		return customErr
	}

	middleware.RespSuccess(ctx, result)

	return nil
}
