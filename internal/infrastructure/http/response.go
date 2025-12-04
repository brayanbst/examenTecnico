package http

type APIResponse[T any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func NewSuccessResponse[T any](message string, data T) APIResponse[T] {
	return APIResponse[T]{
		Code:    "000",
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string) APIResponse[any] {
	return APIResponse[any]{
		Code:    "001",
		Message: message,
	}
}
