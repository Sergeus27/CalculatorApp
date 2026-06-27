package handlers

import "calculator-app/internal/calculationService"

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{
		service: s,
	}
}
