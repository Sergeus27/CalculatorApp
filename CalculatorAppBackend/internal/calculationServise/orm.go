package calculationServise

type Calculation struct { //структура для хранения данных в бд
	ID         string `gorm:"primaryKey" json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct { //структура для запроса на вычисление
	Expression string `json:"expression"`
}
