package services

import (
	api_structure "main/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPublicHoliday interface {
	GetPublicHoliday(filter api_structure.PublicHoliday) ([]api_structure.PublicHoliday, error)
	CreatePublicHoliday(data api_structure.PublicHoliday) ([]api_structure.PublicHoliday, error)
	UpdatePublicHoliday(id int, data api_structure.PublicHoliday) error
	DeletePublicHoliday(id int) error
}

type PublicHolidayService struct{ DB *gorm.DB }

func (r *PublicHolidayService) GetPublicHoliday(filter api_structure.PublicHoliday) ([]api_structure.PublicHoliday, error) {
	result := []api_structure.PublicHoliday{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.PublicHoliday{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *PublicHolidayService) CreatePublicHoliday(data api_structure.PublicHoliday) (api_structure.PublicHoliday, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *PublicHolidayService) UpdatePublicHoliday(Id int, data api_structure.PublicHoliday) error {
	var err error

	if err = r.DB.Model(api_structure.PublicHoliday{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *PublicHolidayService) DeletePublicHoliday(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.PublicHoliday{}).Error; err != nil {
		return err
	}
	return err
}
