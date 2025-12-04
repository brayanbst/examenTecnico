package http

// Response es el formato estándar de respuesta de la API.
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewErrorResponse crea una respuesta de error con código "001".
func NewErrorResponse(message string) Response {
	return Response{
		Code:    "001",
		Message: message,
		Data:    nil,
	}
}

// NewSuccessResponse crea una respuesta de éxito con código "000".
func NewSuccessResponse(message string, data interface{}) Response {
	return Response{
		Code:    "000",
		Message: message,
		Data:    data,
	}
}