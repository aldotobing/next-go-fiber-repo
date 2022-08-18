package responsedto

import (
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

func ErrorResponse(message interface{}) viewmodel.ResponseErrorVM {
	err := []interface{}{message}
	res := viewmodel.ResponseErrorVM{Messages: err}

	return res
}

func ErrorResponseWithCode(msgcode int, message interface{}) viewmodel.ResponseErrorVM {
	err := []interface{}{message}
	type customMessage struct {
		Message      interface{} `json:"err_msg"`
		ResponseCode int         `json:"code"`
	}
	ObjResponse := new(customMessage)
	ObjResponse.ResponseCode = msgcode
	ObjResponse.Message = err
	res := viewmodel.ResponseErrorVM{Messages: ObjResponse}

	return res
}

func SuccessResponse(data interface{}, meta interface{}) viewmodel.ResponseSuccessVM {
	return viewmodel.ResponseSuccessVM{
		Data: data,
		Meta: meta,
	}
}
