package services

import (
	api_structure "main/structures/notifications"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type INotifications interface {
	GetNotifications(filter api_structure.Notifications) ([]api_structure.Notifications, error)
	CreateNotifications(data api_structure.Notifications) ([]api_structure.Notifications, error)
	UpdateNotifications(id int, data api_structure.Notifications) error
	DeleteNotifications(id int) error
}

type NotificationsService struct{ DB *gorm.DB }

func (r *NotificationsService) GetNotifications(filter api_structure.Notifications) ([]api_structure.Notifications, error) {
	result := []api_structure.Notifications{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Notifications{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *NotificationsService) CreateNotifications(data api_structure.Notifications) (api_structure.Notifications, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *NotificationsService) UpdateNotifications(Id int, data api_structure.Notifications) error {
	var err error

	if err = r.DB.Model(api_structure.Notifications{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *NotificationsService) DeleteNotifications(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Notifications{}).Error; err != nil {
		return err
	}
	return err
}
