package httpresponse

type ApiResponse struct {
	StatusCode int
	Msg        string
}

func NewApiError(statusCode int, msg string) *ApiResponse {
	return &ApiResponse{
		StatusCode: statusCode,
		Msg:        msg,
	}
}
