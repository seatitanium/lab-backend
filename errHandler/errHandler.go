package errHandler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type CustomErr struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Request  string `json:"request"`
}

type ApiException struct {
	CustomErr
}

func (e *ApiException) Error() string {
	return e.Msg
}

func (e *CustomErr) Handle() *ApiException {
	return &ApiException{CustomErr{
		HttpCode: e.HttpCode,
		Code:     e.Code,
		Msg:      e.Msg,
	}}
}

func newCustomError(errorType int, httpCode int, errorCode int, errorMsg string) *CustomErr {
	var errMsg4User = errorMsg

	switch errorType {
	case ErrTypeServer:
		logSystemError("Server", errorCode, errorMsg)
		errMsg4User = buildErrMsg(RespErrMsgServerErr, ErrMsgContactAdmin)
	case ErrTypeHttp:
		//Http error should not show system log
	case ErrTypeUser:
		//User error should not show system log
	case ErrTypeDb:
		logSystemError("Database", errorCode, errorMsg)
		errMsg4User = buildErrMsg(RespErrMsgDBErr, ErrMsgContactAdmin)
	case ErrTypeAliyun:
		logSystemError("Aliyun", errorCode, errorMsg)
		errMsg4User = buildErrMsg(RespErrMsgAliyunErr, ErrMsgContactAdmin)
	}

	return &CustomErr{
		HttpCode: httpCode,
		Code:     errorCode,
		Msg:      errMsg4User,
	}
}

func logSystemError(errorType string, errorCode int, errorMsg string) {
	var errorTime = time.Now().Format("06.01.02 15:04:05")
	logStr := fmt.Sprintf("ERROR LOG - %s - [%s %d] Msg: [%s]\n", errorTime, errorType, errorCode, errorMsg)

	color.Red(logStr)
}

func buildErrMsg(strs ...string) string {
	return strings.Join(strs, ", ")
}

func NotFound() *CustomErr {
	return newCustomError(ErrTypeHttp, http.StatusNotFound, RespErrCodeNotFound, RespErrMsgNotFound)
}

func ResNotExist() *CustomErr {
	return newCustomError(ErrTypeUser, http.StatusOK, RespErrCodeResNotExist, RespErrMsgNotExist)
}

// 无效 token 错误。当 token 出现过期等情况时使用。
func InvalidToken() *CustomErr {
	return newCustomError(ErrTypeUser, http.StatusUnauthorized, RespErrCodeInvalidToken, RespErrMsgInvalidToken)
}

// 不良 token 错误。当 token 无法正确解析时使用。
func BadToken() *CustomErr {
	return newCustomError(ErrTypeUser, http.StatusUnauthorized, RespErrCodeBadToken, RespErrMsgBadToken)
}

func ServerError(err error) *CustomErr {
	return newCustomError(ErrTypeServer, http.StatusInternalServerError, RespErrCodeServerErr, err.Error())
}

func AliyunError(err error) *CustomErr {
	return newCustomError(ErrTypeAliyun, http.StatusInternalServerError, RespErrCodeAliyunErr, err.Error())
}

func WrongParam() *CustomErr {
	return newCustomError(ErrTypeUser, http.StatusBadRequest, RespErrCodeParamErr, RespErrMsgParamErr)
}

func UnAuth() *CustomErr {
	return newCustomError(ErrTypeUser, http.StatusUnauthorized, RespErrCodeUnauth, RespErrMsgUnauth)
}

func Forbidden() *CustomErr {
	return newCustomError(ErrTypeHttp, http.StatusForbidden, RespErrCodeForbidden, RespErrMsgForbidden)
}

func OperateNotApplied() *CustomErr {
	return newCustomError(ErrTypeUser, http.StatusOK, RespErrCodeOperateNotApplied, RespErrMsgOperateNotApplied)
}

func DbError(err error) *CustomErr {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ResNotExist()
	}
	return newCustomError(ErrTypeDb, http.StatusInternalServerError, RespErrCodeDBErr, err.Error())
}
