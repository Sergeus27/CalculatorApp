package main

import (
	"fmt"
	//"log"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//var calculations = []Calculation{}

//Основные методы ORM - Create, Find, Update, Delete

func getCalculations(c echo.Context) error {
	var calculations []Calculation

	if err := db.Find(&calculations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calculations"})
	}

	return c.JSON(http.StatusOK, calculations)
	//return c.JSON(http.StatusOK, calculations)
}

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
		Result:     result,
	}

	if err := db.Create(&calc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Coluld not add calculation"})
	}

	return c.JSON(http.StatusCreated, calc)

}

// Функция обновляет существующую запись вычисления в базе данных.
func patchCalculations(c echo.Context) error {
	// 1. Получение id из URL-параметра
	id := c.Param("id") // ← ВОТ ЗДЕСЬ id передается через URL!

	// 2. Парсинг JSON из тела запроса
	var req CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

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
	*/
}

func deleteCalculations(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Calculation{}, "id = ?", id).Error; err != nil {
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

func main() {
	initDB()
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", getCalculations)
	e.POST("/calculations", postCalculations)
	e.PATCH("/calculations/:id", patchCalculations)
	e.DELETE("/calculations/:id", deleteCalculations)

	e.Start("localhost:8080")
}
