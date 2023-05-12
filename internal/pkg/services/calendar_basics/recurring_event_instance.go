package services

import (
	api_structure "main/internal/pkg/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IRecurringEventInstance interface {
	GetRecurringEventInstance(filter api_structure.RecurringEventInstance) ([]api_structure.RecurringEventInstance, error)
	CreateRecurringEventInstance(data api_structure.RecurringEventInstance) ([]api_structure.RecurringEventInstance, error)
	UpdateRecurringEventInstance(id int, data api_structure.RecurringEventInstance) error
	DeleteRecurringEventInstance(id int) error
}

type RecurringEventInstanceService struct{ DB *gorm.DB }

func (r *RecurringEventInstanceService) GetRecurringEventInstance(filter api_structure.RecurringEventInstance) ([]api_structure.RecurringEventInstance, error) {
	result := []api_structure.RecurringEventInstance{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.RecurringEventInstance{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *RecurringEventInstanceService) CreateRecurringEventInstance(data api_structure.RecurringEventInstance) (api_structure.RecurringEventInstance, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *RecurringEventInstanceService) UpdateRecurringEventInstance(Id int, data api_structure.RecurringEventInstance) error {
	var err error

	if err = r.DB.Model(api_structure.RecurringEventInstance{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *RecurringEventInstanceService) DeleteRecurringEventInstance(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.RecurringEventInstance{}).Error; err != nil {
		return err
	}
	return err
}
