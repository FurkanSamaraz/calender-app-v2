package services

import (
	api_structure "main/internal/pkg/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ISpecialHoliday interface {
	GetSpecialHoliday(filter api_structure.SpecialHoliday) ([]api_structure.SpecialHoliday, error)
	CreateSpecialHoliday(data api_structure.SpecialHoliday) ([]api_structure.SpecialHoliday, error)
	UpdateSpecialHoliday(id int, data api_structure.SpecialHoliday) error
	DeleteSpecialHoliday(id int) error
}

type SpecialHolidayService struct{ DB *gorm.DB }

func (r *SpecialHolidayService) GetSpecialHoliday(filter api_structure.SpecialHoliday) ([]api_structure.SpecialHoliday, error) {
	result := []api_structure.SpecialHoliday{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.SpecialHoliday{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *SpecialHolidayService) CreateSpecialHoliday(data api_structure.SpecialHoliday) (api_structure.SpecialHoliday, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *SpecialHolidayService) UpdateSpecialHoliday(Id int, data api_structure.SpecialHoliday) error {
	var err error

	if err = r.DB.Model(api_structure.SpecialHoliday{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *SpecialHolidayService) DeleteSpecialHoliday(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.SpecialHoliday{}).Error; err != nil {
		return err
	}
	return err
}
