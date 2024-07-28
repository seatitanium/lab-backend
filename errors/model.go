package errors

const (
	ErrTypeServer int = iota
	ErrTypeHttp
	ErrTypeUser
	ErrTypeDb
	ErrTypeAliyun
	ErrTypeMinecraft
)

const (
	ErrMsgContactAdmin = "please contact the administrator"
)

// Response codes 11xx for base error
const (
	RespErrCodeParamErr = 1101
	RespErrMsgParamErr  = "Params Error"

	RespErrCodeNotFound = 1102
	RespErrMsgNotFound  = "Not Found"

	RespErrCodeTargetNotExist = 1103
	RespErrMsgTargetNotExist  = "Target not exist"
)

// Response codes 12xx for auth error
const (
	RespErrCodeUnauth = 1201
	RespErrMsgUnauth  = "Authentication Failed"

	RespErrCodeForbidden = 1202
	RespErrMsgForbidden  = "Forbidden"

	RespErrCodeInvalidToken = 1203
	RespErrMsgInvalidToken  = "Invalid token"

	RespErrCodeBadToken = 1204
	RespErrMsgBadToken  = "Bad token"

	RespErrCodeDuplicatedUserRegistration = 1205
	RespErrMsgDuplicatedUserRegistration  = "Duplicated User Registration"

	RespErrCodeDuplicatedMCIDBinding = 1206
	RespErrMsgDuplicatedMCIDBinding  = "Duplicated MCID Binding"
)

// Response code 13xx for server side error
const (
	RespErrCodeServerErr = 1301
	RespErrMsgServerErr  = "Server Error"

	RespErrCodeDBErr = 1302
	RespErrMsgDBErr  = "Database Error"

	RespErrCodeAliyunErr = 1303
	RespErrMsgAliyunErr  = "Aliyun Error"

	RespErrCodeOperationNotApplied = 1304
	RespErrMsgOperationNotApplied  = "Operation Not Applied"
)

const (
	RespErrCodeOffline = 1401
	RespErrMsgOffline  = "Server Offline"
)
