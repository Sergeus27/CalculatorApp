package handlers

import (
	"calculator-app/internal/calculationService" //импорт пакета из которого берем CalculationService
	"github.com/labstack/echo"
	"net/http"
)

/*
создаем структуру хендлера внутри которой будет сервис
внутри репозитуория у нас база данных, внутри сервиса есть репозиторий , а внутри хендлера будет сервис
*/
type CalculationHandler struct {
	service calculationService.CalculationService //обращаемся к пакету: calculationService.* обращаемся к интерфейсу внутри пакета *.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{
		service: s,
	}
}

//Основные методы ORM - Create, Find, Update, Delete

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"}) //?
	}

	return c.JSON(http.StatusOK, calculations) //?

}

func (h *CalculationHandler) PostCalculations(c echo.Context) error {
	var req calculationService.CalculationRequest

	if err := c.Bind(&req); err != nil { //что делает bind?
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	calc, err := h.service.CreateCalculation(req.Expression) //	создаем вычисление calc, идем в service делаем CreateCalculation, берем его из запроса выражение Expression и передаеми в CreateCalculation

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not Create calculation"})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
	}

	return (c.JSON(http.StatusCreated, calc))
	/*
		calc := Calculation{
			ID:         uuid.NewString(),
			Expression: req.Expression,
			Result:     result,
		}

		if err := db.Create(&calc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Coluld not add calculation"})
		}

		return c.JSON(http.StatusCreated, calc)
	*/
}

// Функция обновляет существующую запись вычисления в базе данных.
func (h *CalculationHandler) PatchCalculations(c echo.Context) error {
	// 1. Получение id из URL-параметра
	id := c.Param("id") // ← ВОТ ЗДЕСЬ id передается через URL!

	// 2. Парсинг JSON из тела запроса
	var req calculationService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedCalc, err := h.service.UpdateCalulation(id, req.Expression)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not update calculation"})
	}

	return c.JSON(http.StatusOK, updatedCalc)
	/*
		// 3. Вычисление результата выражения
		result, err := calculateExpression(req.Expression)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid expression"})
		}

		// 4. Поиск записи в БД по id
		var calc Calculation
		if err := db.First(&calc, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not finde expression"})
		}

		// 5. Обновление полей
		calc.Expression = req.Expression
		calc.Result = result

		// 6. Сохранение изменений
		if err := db.Save(&calc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update calculation"})
		}

		// 7. Возврат обновленной записи
		return c.JSON(http.StatusOK, calc)
		/*
				for i, calculation := range calculations {
					if calculation.ID == id {
						calculations[i].Expression = req.Expression
						calculations[i].Result = result
						return c.JSON(http.StatusOK, calculations[i])
					}
				}

			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Calculation not found"})
	*/ //x2
}

func (h *CalculationHandler) DeleteCalculations(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
	/*
		for i, calculation := range calculations {
			if calculation.ID == id {
				calculations = append(calculations[:i], calculations[i+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Calculation not found"})
	*/

}
