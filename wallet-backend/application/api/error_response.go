package api

type errorResponse struct {
	Error string `json:"error"`
}

func getErrorResponse(e error) errorResponse {
	return errorResponse{
		Error: e.Error(),
	}
}
