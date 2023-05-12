package services

import (
	api_structure "main/internal/pkg/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IShift interface {
	GetShift(filter api_structure.Shift) ([]api_structure.Shift, error)
	CreateShift(data api_structure.Shift) ([]api_structure.Shift, error)
	UpdateShift(id int, data api_structure.Shift) error
	DeleteShift(id int) error
}

type ShiftService struct{ DB *gorm.DB }

func (r *ShiftService) GetShift(filter api_structure.Shift) ([]api_structure.Shift, error) {
	result := []api_structure.Shift{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Shift{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *ShiftService) CreateShift(data api_structure.Shift) (api_structure.Shift, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *ShiftService) UpdateShift(Id int, data api_structure.Shift) error {
	var err error

	if err = r.DB.Model(api_structure.Shift{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *ShiftService) DeleteShift(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Shift{}).Error; err != nil {
		return err
	}
	return err
}
