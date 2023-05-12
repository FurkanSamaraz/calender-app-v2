package services

import (
	api_structure "main/internal/pkg/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IExcuseHoliday interface {
	GetExcuseHoliday(filter api_structure.ExcuseHoliday) ([]api_structure.ExcuseHoliday, error)
	CreateExcuseHoliday(data api_structure.ExcuseHoliday) ([]api_structure.ExcuseHoliday, error)
	UpdateExcuseHoliday(id int, data api_structure.ExcuseHoliday) error
	DeleteExcuseHoliday(id int) error
}

type ExcuseHolidayService struct{ DB *gorm.DB }

func (r *ExcuseHolidayService) GetExcuseHoliday(filter api_structure.ExcuseHoliday) ([]api_structure.ExcuseHoliday, error) {
	result := []api_structure.ExcuseHoliday{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.ExcuseHoliday{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *ExcuseHolidayService) CreateExcuseHoliday(data api_structure.ExcuseHoliday) (api_structure.ExcuseHoliday, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *ExcuseHolidayService) UpdateExcuseHoliday(Id int, data api_structure.ExcuseHoliday) error {
	var err error

	if err = r.DB.Model(api_structure.ExcuseHoliday{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *ExcuseHolidayService) DeleteExcuseHoliday(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.ExcuseHoliday{}).Error; err != nil {
		return err
	}
	return err
}
