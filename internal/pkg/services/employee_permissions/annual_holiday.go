package services

import (
	api_structure "main/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAnnualHoliday interface {
	GetAnnualHoliday(filter api_structure.AnnualHoliday) ([]api_structure.AnnualHoliday, error)
	CreateAnnualHoliday(data api_structure.AnnualHoliday) ([]api_structure.AnnualHoliday, error)
	UpdateAnnualHoliday(id int, data api_structure.AnnualHoliday) error
	DeleteAnnualHoliday(id int) error
}

type AnnualHolidayService struct{ DB *gorm.DB }

func (r *AnnualHolidayService) GetAnnualHoliday(filter api_structure.AnnualHoliday) ([]api_structure.AnnualHoliday, error) {
	result := []api_structure.AnnualHoliday{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.AnnualHoliday{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *AnnualHolidayService) CreateAnnualHoliday(data api_structure.AnnualHoliday) (api_structure.AnnualHoliday, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *AnnualHolidayService) UpdateAnnualHoliday(Id int, data api_structure.AnnualHoliday) error {
	var err error

	if err = r.DB.Model(api_structure.AnnualHoliday{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *AnnualHolidayService) DeleteAnnualHoliday(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.AnnualHoliday{}).Error; err != nil {
		return err
	}
	return err
}
