package services

import (
	api_structure "main/internal/pkg/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IEventInstance interface {
	GetEventInstance(filter api_structure.EventInstance) ([]api_structure.EventInstance, error)
	CreateEventInstance(data api_structure.EventInstance) ([]api_structure.EventInstance, error)
	UpdateEventInstance(id int, data api_structure.EventInstance) error
	DeleteEventInstance(id int) error
}

type EventInstanceService struct{ DB *gorm.DB }

func (r *EventInstanceService) GetEventInstance(filter api_structure.EventInstance) ([]api_structure.EventInstance, error) {
	result := []api_structure.EventInstance{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.EventInstance{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *EventInstanceService) CreateEventInstance(data api_structure.EventInstance) (api_structure.EventInstance, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *EventInstanceService) UpdateEventInstance(Id int, data api_structure.EventInstance) error {
	var err error

	if err = r.DB.Model(api_structure.EventInstance{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *EventInstanceService) DeleteEventInstance(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.EventInstance{}).Error; err != nil {
		return err
	}
	return err
}
