package services

import (
	api_structure "main/internal/pkg/structures/employee_requests"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IEmployeeDevelopment interface {
	GetEmployeeDevelopment(filter api_structure.EmployeeDevelopment) ([]api_structure.EmployeeDevelopment, error)
	CreateEmployeeDevelopment(data api_structure.EmployeeDevelopment) ([]api_structure.EmployeeDevelopment, error)
	UpdateEmployeeDevelopment(id int, data api_structure.EmployeeDevelopment) error
	DeleteEmployeeDevelopment(id int) error
}

type EmployeeDevelopmentService struct{ DB *gorm.DB }

func (r *EmployeeDevelopmentService) GetEmployeeDevelopment(filter api_structure.EmployeeDevelopment) ([]api_structure.EmployeeDevelopment, error) {
	result := []api_structure.EmployeeDevelopment{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.EmployeeDevelopment{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *EmployeeDevelopmentService) CreateEmployeeDevelopment(data api_structure.EmployeeDevelopment) (api_structure.EmployeeDevelopment, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *EmployeeDevelopmentService) UpdateEmployeeDevelopment(Id int, data api_structure.EmployeeDevelopment) error {
	var err error

	if err = r.DB.Model(api_structure.EmployeeDevelopment{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *EmployeeDevelopmentService) DeleteEmployeeDevelopment(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.EmployeeDevelopment{}).Error; err != nil {
		return err
	}
	return err
}
