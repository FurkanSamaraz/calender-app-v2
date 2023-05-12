package services

import (
	api_structure "main/internal/pkg/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IRecurringEvent interface {
	GetRecurringEvent(filter api_structure.RecurringEvent) ([]api_structure.RecurringEvent, error)
	CreateRecurringEvent(data api_structure.RecurringEvent) ([]api_structure.RecurringEvent, error)
	UpdateRecurringEvent(id int, data api_structure.RecurringEvent) error
	DeleteRecurringEvent(id int) error
}

type RecurringEventService struct{ DB *gorm.DB }

func (r *RecurringEventService) GetRecurringEvent(filter api_structure.RecurringEvent) ([]api_structure.RecurringEvent, error) {
	result := []api_structure.RecurringEvent{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.RecurringEvent{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *RecurringEventService) CreateRecurringEvent(data api_structure.RecurringEvent) (api_structure.RecurringEvent, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *RecurringEventService) UpdateRecurringEvent(Id int, data api_structure.RecurringEvent) error {
	var err error

	if err = r.DB.Model(api_structure.RecurringEvent{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *RecurringEventService) DeleteRecurringEvent(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.RecurringEvent{}).Error; err != nil {
		return err
	}
	return err
}
