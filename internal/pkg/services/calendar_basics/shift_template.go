package services

import (
	api_structure "main/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IShiftTemplate interface {
	GetShiftTemplate(filter api_structure.ShiftTemplate) ([]api_structure.ShiftTemplate, error)
	CreateShiftTemplate(data api_structure.ShiftTemplate) ([]api_structure.ShiftTemplate, error)
	UpdateShiftTemplate(id int, data api_structure.ShiftTemplate) error
	DeleteShiftTemplate(id int) error
}

type ShiftTemplateService struct{ DB *gorm.DB }

func (r *ShiftTemplateService) GetShiftTemplate(filter api_structure.ShiftTemplate) ([]api_structure.ShiftTemplate, error) {
	result := []api_structure.ShiftTemplate{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.ShiftTemplate{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *ShiftTemplateService) CreateShiftTemplate(data api_structure.ShiftTemplate) (api_structure.ShiftTemplate, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *ShiftTemplateService) UpdateShiftTemplate(Id int, data api_structure.ShiftTemplate) error {
	var err error

	if err = r.DB.Model(api_structure.ShiftTemplate{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *ShiftTemplateService) DeleteShiftTemplate(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.ShiftTemplate{}).Error; err != nil {
		return err
	}
	return err
}
