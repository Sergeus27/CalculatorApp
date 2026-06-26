package calculationService

import (
	"gorm.io/gorm"
)

//Основные методы CRUD- Create, Read. Update, Delete

// интерфейс для работы с базой данных
type CalculationRepository interface {
	CreateCalculation(calc Calculation) error          //создает вычисление в базу данных
	GetAllCalculation() ([]Calculation, error)         //получаем слайс всех вычислений из БД
	GetCalculationByID(id string) (Calculation, error) //поиск выражения по его id
	UpdateCalulation(calc Calculation) error           //обновить результат
	DeleteCalculation(id string) error                 //удалить вчисление по id
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}
}

func (r *calcRepository) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *calcRepository) GetAllCalculation() ([]Calculation, error) {
	var calculations []Calculation
	err := r.db.Find(&calculations).Error

	return calculations, err
}

func (r *calcRepository) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation
	err := r.db.First(&calc, "id = ?", id).Error

	return calc, err
}

func (r *calcRepository) UpdateCalulation(calc Calculation) error {
	return r.db.Save(&calc).Error
}

func (r *calcRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
