package services

import (
	api_structure "main/internal/pkg/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBirthHoliday interface {
	GetBirthHoliday(filter api_structure.BirthHoliday) ([]api_structure.BirthHoliday, error)
	CreateBirthHoliday(data api_structure.BirthHoliday) ([]api_structure.BirthHoliday, error)
	UpdateBirthHoliday(id int, data api_structure.BirthHoliday) error
	DeleteBirthHoliday(id int) error
}

type BirthHolidayService struct{ DB *gorm.DB }

func (r *BirthHolidayService) GetBirthHoliday(filter api_structure.BirthHoliday) ([]api_structure.BirthHoliday, error) {
	result := []api_structure.BirthHoliday{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.BirthHoliday{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *BirthHolidayService) CreateBirthHoliday(data api_structure.BirthHoliday) (api_structure.BirthHoliday, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *BirthHolidayService) UpdateBirthHoliday(Id int, data api_structure.BirthHoliday) error {
	var err error

	if err = r.DB.Model(api_structure.BirthHoliday{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *BirthHolidayService) DeleteBirthHoliday(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.BirthHoliday{}).Error; err != nil {
		return err
	}
	return err
}
