package service

import (
	"fmt"
	api_structure_chat "main/internal/pkg/structures/chat"
	api_structure "main/internal/pkg/structures/employee_requests"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatService struct{ DB *gorm.DB }
type IChatInstance interface {
	Login(filter api_structure.Employee) (string, error)
	Register(data api_structure.Employee) (api_structure_chat.StatusMessage, error)

	VerifyContact(username string) *api_structure_chat.Response
	ChatHistory(username1, username2, fromTS, toTS string) *api_structure_chat.Response
	ContactList(username string) *api_structure_chat.Response
}

func (r *ChatService) Login(filter api_structure.Employee) (string, error) {
	result := api_structure.Employee{}

	var err error

	if err = r.DB.Table(result.TableName()).Preload(clause.Associations).Model(&api_structure.Employee{}).Where(filter).Find(&result).Error; err != nil {
		fmt.Printf("not user error")
		return err.Error(), err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Name,
		"password": result.Password,
	})
	tokenString, err := token.SignedString([]byte("gizli-anahtar"))
	if err != nil {
		fmt.Println("hatalı token")
	}
	return tokenString, err
}

func (r *ChatService) Register(data api_structure.Employee) (api_structure_chat.StatusMessage, error) {
	result := api_structure_chat.StatusMessage{}
	var err error
	if err = r.DB.Table(result.TableName()).Create(&data).Error; err != nil {
		result.Message = "Error Register"
		return result, err
	}
	result.Message = "Successfully Register"

	return result, err
}
