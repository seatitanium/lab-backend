package bss

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/aliyun/bss"
	"seatimc/backend/errors"
	"seatimc/backend/middleware"
	"strconv"
)

func HandleQueryAccountTransactions(ctx *gin.Context) *errors.CustomErr {
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
			return errors.WrongParam()
		}

		if _pagenumN <= 0 {
			return errors.WrongParam()
		}
	}

	if pagesize == "" {
		_pagesizeN = 10
	} else {
		_pagesizeN, err = strconv.ParseInt(pagesize, 10, 32)

		if err != nil {
			return errors.WrongParam()
		}

		if _pagesizeN <= 0 {
			return errors.WrongParam()
		}
	}

	var pagenumN = int32(_pagenumN)
	var pagesizeN = int32(_pagesizeN)

	result, customErr := bss.QueryAccountTransactions(pagenumN, pagesizeN)

	if err != nil {
		return customErr
	}

	middleware.RespSuccess(ctx, result)

	return nil
}
