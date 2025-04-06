package common

type ErrorData struct {
	Error string `json:"error"`
}
type CustomResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func CustomRespCreator(data any, err error) CustomResponse {
	if err != nil {
		return CustomResponse{Success: false, Data: ErrorData{Error: err.Error()}}
	} else {
		return CustomResponse{Success: true, Data: data}
	}
}
