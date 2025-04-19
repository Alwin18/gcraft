package utils

type Meta struct {
	Page      int `json:"page" example:"1"`
	PerPage   int `json:"per_page" example:"10"`
	TotalData int `json:"total_data" example:"100"`
	Totalpage int `json:"total_page" example:"10"`
}

type ProcessResponse struct {
	Message string `json:"message" example:"succes get data"`
	Code    int    `json:"code" example:"200"`
}

type ResponseError struct {
	Message string   `json:"message"`
	Code    int      `json:"code"`
	Errors  []string `json:"errors,omitempty"`
}

type Response[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message" example:"succes get data"`
	Code    int    `json:"code" example:"200"`
}

type ResponseWithPagination[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message" example:"succes get data"`
	Code    int    `json:"code" example:"200"`
	Meta    Meta   `json:"meta"`
}

func NewResponse[T any](data T, message string, code int) Response[T] {
	return Response[T]{
		Data:    data,
		Message: message,
		Code:    code,
	}
}

func NewResponseWithPagination[T any](data T, message string, code int, meta Meta) ResponseWithPagination[T] {
	return ResponseWithPagination[T]{
		Data:    data,
		Message: message,
		Code:    code,
		Meta:    meta,
	}
}

func NewErrorResponse(err ResponseError) ResponseError {
	return ResponseError{
		Message: err.Message,
		Code:    err.Code,
		Errors:  err.Errors,
	}
}

func NewProcessResponse(message string, code int) ProcessResponse {
	return ProcessResponse{
		Message: message,
		Code:    code,
	}
}
