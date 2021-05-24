package serializer

import (
	"go-skeleton/utils/errors"
	"reflect"
)

type Response struct {
	Code    int         `json:"code"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

func NewResponse(code int, msg string, result interface{}) Response {
	return Response{
		Code:    code,
		Result:  result,
		Message: msg,
	}
}

func ReplyError(err error) Response {
	_, hasMethod := reflect.TypeOf(err).MethodByName("Code")
	if hasMethod {
		err2 := err.(*errors.Err)
		return Response{Code: err2.Code(), Result: "", Message: err.Error()}
	}
	return Response{Code: errors.ERROR, Result: "", Message: err.Error()}
}

func ReplySuccess(result interface{}) Response {
	return Response{Code: errors.SUCCESS, Result: result, Message: errors.GetErrMsg(errors.SUCCESS)}
}
