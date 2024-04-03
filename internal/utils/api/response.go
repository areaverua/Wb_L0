package response

import "github.com/areaverua/l0_task/internal/repository/models"

type Response struct {
	Status string       `json:"status"`
	Error  string       `json:"error,omitempty"`
	Data   models.Order `json:"order,omitempty"'`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK(order models.Order) Response {
	return Response{
		Status: StatusOK,
		Data:   order,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
