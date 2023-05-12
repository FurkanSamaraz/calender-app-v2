package services

import (
	"fmt"
	Employees_api_structure "main/structures/employee_requests"
	api_structure "main/structures/notifications"

	"github.com/go-gomail/gomail"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ISend_Emails interface {
	GetSend_Emails(filter api_structure.Send_Emails) ([]api_structure.Send_Emails, error)
	CreateSend_Emails(data api_structure.Send_Emails) ([]api_structure.Send_Emails, error)
	UpdateSend_Emails(id int, data api_structure.Send_Emails) error
	DeleteSend_Emails(id int) error
	Send_SendEmail(notification api_structure.Send_Emails) error
}

type Send_EmailsService struct{ DB *gorm.DB }

func (r *Send_EmailsService) GetSend_Emails(filter api_structure.Send_Emails) ([]api_structure.Send_Emails, error) {
	result := []api_structure.Send_Emails{}
	var err error

	if err = r.DB.Preload(clause.Associations).Model(&api_structure.Send_Emails{}).Where(filter).Find(&result).Error; err != nil {
		return result, err
	}
	return result, err
}
func (r *Send_EmailsService) CreateSend_Emails(data api_structure.Send_Emails) (api_structure.Send_Emails, error) {
	var err error
	if err = r.DB.Create(&data).Error; err != nil {
		return data, err
	}
	return data, err
}
func (r *Send_EmailsService) UpdateSend_Emails(Id int, data api_structure.Send_Emails) error {
	var err error

	if err = r.DB.Model(api_structure.Send_Emails{}).Where("id = ?", Id).Updates(&data).Error; err != nil {
		return err
	}
	return err
}
func (r *Send_EmailsService) DeleteSend_Emails(Id int) error {
	var err error
	if err = r.DB.Where("id = ?", Id).Delete(&api_structure.Send_Emails{}).Error; err != nil {
		return err
	}
	return err
}

func (r *Send_EmailsService) Send_SendEmail(Id int) error {
	notification := api_structure.Send_Emails{}
	if err := r.DB.Preload(clause.Associations).Model(&api_structure.Send_Emails{}).Where("id = ?", Id).First(&notification).Error; err != nil {
		return err
	}

	result := Employees_api_structure.Employee{}
	if err := r.DB.Preload(clause.Associations).Model(&Employees_api_structure.Employee{}).Where("id = ?", notification.EmployeeID).First(&result).Error; err != nil {
		return err
	}

	// E-posta nesnesini oluştur
	m := gomail.NewMessage()
	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", result.Email)
	m.SetHeader("Subject", notification.Subject)
	m.SetBody("text/plain", notification.Body)
	// SMTP sunucusuna bağlan
	d := gomail.NewDialer("smtp.example.com", 587, "sender@example.com", "password")
	// E-postayı gönder
	err := d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	return err
}
