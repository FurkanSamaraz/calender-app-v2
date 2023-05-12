package services

import (
	api_structure "main/internal/pkg/structures/calendar_basics"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IImportantDay interface {
	GetImportantDay(filter api_structure.ImportantDay) ([]api_structure.ImportantDay, error)
	CreateImportantDay(data api_structure.ImportantDay) ([]api_structure.ImportantDay, error)
	UpdateImportantDay(id int, data api_structure.ImportantDay) error
	DeleteImportantDay(id int) error
}

type ImportantDayService struct{ DB *gorm.DB }

func (r *ImportantDayService) GetImportantDay(filter api_structure.ImportantDay) ([]api_structure.ImportantDay, error) {
	result := []api_structure.ImportantDay{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.ImportantDay{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *ImportantDayService) CreateImportantDay(data api_structure.ImportantDay) (api_structure.ImportantDay, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *ImportantDayService) UpdateImportantDay(Id int, data api_structure.ImportantDay) error {
	var err error

	if err = r.DB.Model(api_structure.ImportantDay{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *ImportantDayService) DeleteImportantDay(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.ImportantDay{}).Error; err != nil {
		return err
	}
	return err
}
