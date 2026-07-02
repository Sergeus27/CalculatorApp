package main

import (
	"calculator-app/internal/calculationService"
	"calculator-app/internal/db"
	"calculator-app/internal/handlers"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//var calculations = []Calculation{}

func main() {
	//создаем экземпляры структуры репозитория сервиса и хендлера
	//в репозиторий надо передать базу данных в сервис репизиторий в хендлер сервис

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}
