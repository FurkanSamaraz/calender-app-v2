package services

import (
	api_structure "main/internal/pkg/structures/employee_requests"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IEmployeeEventRequest interface {
	GetEmployeeEventRequest(filter api_structure.EmployeeEventRequest) ([]api_structure.EmployeeEventRequest, error)
	CreateEmployeeEventRequest(data api_structure.EmployeeEventRequest) ([]api_structure.EmployeeEventRequest, error)
	UpdateEmployeeEventRequest(id int, data api_structure.EmployeeEventRequest) error
	DeleteEmployeeEventRequest(id int) error
}

type EmployeeEventRequestService struct{ DB *gorm.DB }

func (r *EmployeeEventRequestService) GetEmployeeEventRequest(filter api_structure.EmployeeEventRequest) ([]api_structure.EmployeeEventRequest, error) {
	result := []api_structure.EmployeeEventRequest{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.EmployeeEventRequest{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *EmployeeEventRequestService) CreateEmployeeEventRequest(data api_structure.EmployeeEventRequest) (api_structure.EmployeeEventRequest, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *EmployeeEventRequestService) UpdateEmployeeEventRequest(Id int, data api_structure.EmployeeEventRequest) error {
	var err error

	if err = r.DB.Model(api_structure.EmployeeEventRequest{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *EmployeeEventRequestService) DeleteEmployeeEventRequest(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.EmployeeEventRequest{}).Error; err != nil {
		return err
	}
	return err
}
