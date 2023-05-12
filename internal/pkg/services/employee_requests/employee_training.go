package services

import (
	api_structure "main/internal/pkg/structures/employee_requests"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IEmployeeTraining interface {
	GetEmployeeTraining(filter api_structure.EmployeeTraining) ([]api_structure.EmployeeTraining, error)
	CreateEmployeeTraining(data api_structure.EmployeeTraining) ([]api_structure.EmployeeTraining, error)
	UpdateEmployeeTraining(id int, data api_structure.EmployeeTraining) error
	DeleteEmployeeTraining(id int) error
}

type EmployeeTrainingService struct{ DB *gorm.DB }

func (r *EmployeeTrainingService) GetEmployeeTraining(filter api_structure.EmployeeTraining) ([]api_structure.EmployeeTraining, error) {
	result := []api_structure.EmployeeTraining{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.EmployeeTraining{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *EmployeeTrainingService) CreateEmployeeTraining(data api_structure.EmployeeTraining) (api_structure.EmployeeTraining, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *EmployeeTrainingService) UpdateEmployeeTraining(Id int, data api_structure.EmployeeTraining) error {
	var err error

	if err = r.DB.Model(api_structure.EmployeeTraining{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *EmployeeTrainingService) DeleteEmployeeTraining(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.EmployeeTraining{}).Error; err != nil {
		return err
	}
	return err
}
