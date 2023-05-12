package services

import (
	api_structure "main/structures/employee_permissions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IPermission interface {
	GetPermission(filter api_structure.Permission) ([]api_structure.Permission, error)
	CreatePermission(data api_structure.Permission) ([]api_structure.Permission, error)
	UpdatePermission(id int, data api_structure.Permission) error
	DeletePermission(id int) error
}

type PermissionService struct{ DB *gorm.DB }

func (r *PermissionService) GetPermission(filter api_structure.Permission) ([]api_structure.Permission, error) {
	result := []api_structure.Permission{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Permission{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *PermissionService) CreatePermission(data api_structure.Permission) (api_structure.Permission, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *PermissionService) UpdatePermission(Id int, data api_structure.Permission) error {
	var err error

	if err = r.DB.Model(api_structure.Permission{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *PermissionService) DeletePermission(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Permission{}).Error; err != nil {
		return err
	}
	return err
}
