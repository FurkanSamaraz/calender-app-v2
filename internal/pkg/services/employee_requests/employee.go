package services

import (
	api_structure "main/structures/employee_requests"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IEmployee interface {
	GetEmployee(filter api_structure.Employee) ([]api_structure.Employee, error)
	CreateEmployee(data api_structure.Employee) ([]api_structure.Employee, error)
	UpdateEmployee(id int, data api_structure.Employee) error
	DeleteEmployee(id int) error
}

type EmployeeService struct{ DB *gorm.DB }

func (r *EmployeeService) GetEmployee(filter api_structure.Employee) ([]api_structure.Employee, error) {
	result := []api_structure.Employee{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Employee{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *EmployeeService) CreateEmployee(data api_structure.Employee) (api_structure.Employee, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *EmployeeService) UpdateEmployee(Id int, data api_structure.Employee) error {
	var err error

	if err = r.DB.Model(api_structure.Employee{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *EmployeeService) DeleteEmployee(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Employee{}).Error; err != nil {
		return err
	}
	return err
}
