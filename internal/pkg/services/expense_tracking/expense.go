package services

import (
	api_structure "main/structures/expense_tracking"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IExpense interface {
	GetExpense(filter api_structure.Expense) ([]api_structure.Expense, error)
	CreateExpense(data api_structure.Expense) ([]api_structure.Expense, error)
	UpdateExpense(id int, data api_structure.Expense) error
	DeleteExpense(id int) error
}

type ExpenseService struct{ DB *gorm.DB }

func (r *ExpenseService) GetExpense(filter api_structure.Expense) ([]api_structure.Expense, error) {
	result := []api_structure.Expense{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Expense{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *ExpenseService) CreateExpense(data api_structure.Expense) (api_structure.Expense, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *ExpenseService) UpdateExpense(Id int, data api_structure.Expense) error {
	var err error

	if err = r.DB.Model(api_structure.Expense{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *ExpenseService) DeleteExpense(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Expense{}).Error; err != nil {
		return err
	}
	return err
}
