package service

import (
	"encoding/json"
	"github.com/areaverua/l0_task/internal/repository/models"
	"github.com/go-playground/validator/v10"
	"log/slog"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) OrderIsValid(data []byte, log *slog.Logger) (bool, *models.OrderDTO) {
	var order models.OrderDTO
	err := json.Unmarshal(data, &order)
	if err != nil {
		log.Info("Validation: message decoding fail", err)
		return false, nil
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(order)
	if err != nil {
		log.Info("Validation: message is invalid: ", err)
		return false, nil
	}
	return true, &order
}
