package services

import (
	api_structure "main/internal/pkg/structures/integration"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IAuthProvider interface {
	GetAuthProvider(filter api_structure.AuthProvider) ([]api_structure.AuthProvider, error)
	CreateAuthProvider(data api_structure.AuthProvider) ([]api_structure.AuthProvider, error)
	UpdateAuthProvider(id int, data api_structure.AuthProvider) error
	DeleteAuthProvider(id int) error
}

type AuthProviderService struct{ DB *gorm.DB }

func (r *AuthProviderService) GetAuthProvider(filter api_structure.AuthProvider) ([]api_structure.AuthProvider, error) {
	result := []api_structure.AuthProvider{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.AuthProvider{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *AuthProviderService) CreateAuthProvider(data api_structure.AuthProvider) (api_structure.AuthProvider, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *AuthProviderService) UpdateAuthProvider(Id int, data api_structure.AuthProvider) error {
	var err error

	if err = r.DB.Model(api_structure.AuthProvider{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *AuthProviderService) DeleteAuthProvider(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.AuthProvider{}).Error; err != nil {
		return err
	}
	return err
}
