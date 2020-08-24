package response

import "net/http"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"-"`
}

func NewCreatedResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
		Status:  http.StatusCreated,
	}
}

func NewOkResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
		Status:  http.StatusOK,
	}
}

func NewNotFoundError(message string) *BaseResponse {
	return &BaseResponse{
		Success: false,
		Message: message,
		Status:  http.StatusNotFound,
		Data:    nil,
	}
}

func NewBadRequestError(message string) *BaseResponse {
	return &BaseResponse{
		Success: false,
		Status:  http.StatusBadRequest,
		Message: message,
		Data:    nil,
	}
}

func NewInternalServerError(message string) *BaseResponse {
	return &BaseResponse{
		Success: false,
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}

func NewUnAuthorizedError() *BaseResponse {
	return &BaseResponse{
		Success: false,
		Status:  http.StatusUnauthorized,
		Message: Unauthorized,
		Data:    nil,
	}
}
