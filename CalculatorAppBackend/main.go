AXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXpackage main

import (
	"fmt"
	"net/http"анных в бд
	ID         string `json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct { //структура для запроса на вычисление
	Expression string `json:"expression"`
}

var calculations = []Calculation{}

func calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)

	if err != nil {
		return "", err //ошибка 2+2a
	}

AXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAX	result, err := expr.Evaluate(nil)

	if err != nil {
		ret\ult), err
}

func getCalculations(c error {
	return c.JSON(http.StatusOK, calculations)
AXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAX}
ваыцавыаывавыаыва
func postCalculations(c echo.Context) error {
	var req CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	result, err := calculateExpression(req.Expression)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}
	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
AXXAXAXAXXAXAXAXXAXAXAXXAXAX		Result:     result,
	}

	calculations = append(calculations, calc)
	return c.JSON(http.StatusCreated, calc)
AXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAX
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)
	e.Start("localhost:8080")
}
AXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAXAXXAXAX