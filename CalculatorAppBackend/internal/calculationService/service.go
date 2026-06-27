package calculationService

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

// посредник между нашим сайтом и базой даннных, здесь происходят вычисления
type CalculationService interface {
	CreateCalculation(string) (Calculation, error) //измеение здесь передаем строку а не структуру
	GetAllCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalulation(id, expression string) (Calculation, error) // передаем из ручки новое выражение и возвращаем готовое посчитанное
	DeleteCalculation(id string) error
}

type calcService struct {
	repo CalculationRepository
}

func NewCalculationService(r CalculationRepository) CalculationService {
	return &calcService{repo: r} //тут мне не понятно разве 2 интерфейса с отличающимися методами не конфликтуют?
}

func (s *calcService) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)

	if err != nil {
		return "", err //ошибка 2+2a
	}

	result, err := expr.Evaluate(nil)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}

// создать вычисление для базы данных
func (s *calcService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculateExpression(expression) //во вспомогательной функции считаем то что передали нам из ручки
	if err != nil {
		return Calculation{}, err //если произошла ошибка на подсчете возвращаем пустое вычисление в бащу данных не идем
	}
	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     result,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, nil

}

// получить все выражения
func (s *calcService) GetAllCalculation() ([]Calculation, error) {
	return s.repo.GetAllCalculation() //обращаемся в наш репозиторий и получаем все выражения
}

// получаем выражение по id
func (s *calcService) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

// UpdateCalulation implements calculationServise.
func (s *calcService) UpdateCalulation(id string, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}
	calc.Expression = expression
	calc.Result = result

	if err := s.repo.UpdateCalulation(calc); err != nil {
		return Calculation{}, err
	}
	return calc, nil
}

func (s *calcService) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}
