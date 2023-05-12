package services

import (
	api_structure "main/internal/pkg/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IEvent interface {
	GetEvent(filter api_structure.Event) ([]api_structure.Event, error)
	CreateEvent(data api_structure.Event) ([]api_structure.Event, error)
	UpdateEvent(id int, data api_structure.Event) error
	DeleteEvent(id int) error
}

type EventService struct{ DB *gorm.DB }

func (r *EventService) GetEvent(filter api_structure.Event) ([]api_structure.Event, error) {
	result := []api_structure.Event{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Event{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *EventService) CreateEvent(data api_structure.Event) (api_structure.Event, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *EventService) UpdateEvent(Id int, data api_structure.Event) error {
	var err error

	if err = r.DB.Model(api_structure.Event{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *EventService) DeleteEvent(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Event{}).Error; err != nil {
		return err
	}
	return err
}
