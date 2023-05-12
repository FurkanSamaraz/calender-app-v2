package services

import (
	api_structure "main/internal/pkg/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAdministrativeHoliday interface {
	GetAdministrativeHoliday(filter api_structure.AdministrativeHoliday) ([]api_structure.AdministrativeHoliday, error)
	CreateAdministrativeHoliday(data api_structure.AdministrativeHoliday) ([]api_structure.AdministrativeHoliday, error)
	UpdateAdministrativeHoliday(id int, data api_structure.AdministrativeHoliday) error
	DeleteAdministrativeHoliday(id int) error
}

type AdministrativeHolidayService struct{ DB *gorm.DB }

func (r *AdministrativeHolidayService) GetAdministrativeHoliday(filter api_structure.AdministrativeHoliday) ([]api_structure.AdministrativeHoliday, error) {
	result := []api_structure.AdministrativeHoliday{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.AdministrativeHoliday{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *AdministrativeHolidayService) CreateAdministrativeHoliday(data api_structure.AdministrativeHoliday) (api_structure.AdministrativeHoliday, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *AdministrativeHolidayService) UpdateAdministrativeHoliday(Id int, data api_structure.AdministrativeHoliday) error {
	var err error

	if err = r.DB.Model(api_structure.AdministrativeHoliday{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *AdministrativeHolidayService) DeleteAdministrativeHoliday(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.AdministrativeHoliday{}).Error; err != nil {
		return err
	}
	return err
}
