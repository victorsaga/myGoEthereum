package CommonModel

import (
	"myGoEthereum/Model/ResultCode"
)

type ApiResponseWithData struct {
	ApiResponse
	Data interface{} `json:"data"`
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (a *ApiResponse) SetFalseUnknow(message string) {
	a.Success = false
	a.Code = ResultCode.Unknown
	a.Message = message
}

func (a *ApiResponse) SetError(code string, message string) {
	a.Success = false
	a.Code = code
	a.Message = message
}

func (a *ApiResponse) SetSuccess() {
	a.Success = true
	a.Code = ResultCode.Success
	a.Message = ""
}
func (a *ApiResponse) SetSuccessMessage(message string) {
	a.Success = true
	a.Code = ResultCode.Success
	a.Message = message
}

type CommonLogContent struct {
	Type           string
	LogId          string
	Route          string
	HttpMethod     string
	HttpStatusCode int
	Header         interface{}
	Request        interface{}
	Response       interface{}
	Error          interface{}
	ExecuteSecond  float64
	RequestTime    string
	ResponseTime   string
	LogTimeZone    string
}
