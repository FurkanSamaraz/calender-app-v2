package controller

import (
	"fmt"

	model "main/internal/pkg/structures/chat"

	api_service "main/internal/pkg/service/chat"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	Svc api_service.ChatService
}

// RegisterHandler godoc
// @Summary       Register User
// @Description   Registers a new user
// @Tags          User
// @Accept        json
// @Produce       json
// @Param         body body model.User true "Request body"
// @Success       200 {object} model.Response
// @Failure       500 {object} model.ErrorResponse
// @Router        /register [post]
func (controller *ChatController) RegisterHandler(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Invalid request",
		})
	}
	res := Register(&user)
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Username and password are required",
		})
	}
	result, rerr := controller.Svc.Register(user)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Fetch Data",
			Message: rerr.Error(),
		})
	}
	fmt.Println(result)
	res.Message = result.Message

	return c.Status(fiber.StatusOK).JSON(res)
}

// LoginHandler godoc
// @Summary       Login User
// @Description   Logs in a user
// @Tags          User
// @Accept        json
// @Produce       json
// @Param         body body model.User true "Request body"
// @Success       200 {object} model.Response
// @Failure       500 {object} model.ErrorResponse
// @Router        /login [post]
func (controller *ChatController) LoginHandler(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Invalid request",
		})
	}

	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Fetch Data",
			Message: "Username and password are required",
		})
	}

	result, rerr := controller.Svc.Login(user)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Create Data",
			Message: rerr.Error(),
		})
	}
	c.Locals("Authorization", result)

	res := Login(&user)
	res.Jwt = result

	return c.Status(fiber.StatusOK).JSON(res)
}

// VerifyContactHandler godoc
// @Summary       Verify Contact
// @Description   Verifies the contact of a user
// @Tags          User
// @Security      BearerAuth
// @Accept        json
// @Produce       json
// @Success       200 {object} model.Response
// @Failure       500 {object} model.ErrorResponse
// @Router        /verify-contact [get]
func (controller *ChatController) VerifyContactHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	res, err := controller.Svc.VerifyContact(claims["username"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Contact Data",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}

// ChatHistoryHandler godoc
// @Summary       Chat History
// @Description   Retrieves the chat history between two users
// @Tags          Chat
// @Accept        json
// @Produce       json
// @Param         u1 query string true "Username of user 1"
// @Param         u2 query string true "Username of user 2"
// @Success       200 {object} model.Response
// @Failure       500 {object} model.ErrorResponse
// @Router        /chat-history [get]
func (controller *ChatController) ChatHistoryHandler(c *fiber.Ctx) error {
	username1 := c.Query("u1")
	username2 := c.Query("u2")

	res, err := controller.Svc.ChatHistory(username1, username2)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Chat History Data",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}

// ContactListHandler godoc
// @Summary       Contact List
// @Description   Retrieves the contact list of a user
// @Tags          User
// @Security      BearerAuth
// @Accept        json
// @Produce       json
// @Success       200 {object} model.ContactList
// @Failure       500 {object} model.ErrorResponse
// @Router        /contact-list [get]
func (controller *ChatController) ContactListHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	res, err := controller.Svc.ContactList(claims["username"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponse{
			Type:    "Chat List Data",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}

// AuthMiddleware godoc
// @Summary       Authentication Middleware
// @Description   Authenticates the request using JWT token
// @Tags          Authentication
// @Security      ApiKeyAuth
// @Accept        json
// @Produce       json
// @Success       200 {object} model.TokenClaims
// @Failure       401 {object} model.AuthErrorResponse
// @Failure       500 {object} model.AuthErrorResponse
// @Router        /auth-middleware [get]
func (controller *ChatController) AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Eksik yetkilendirme belirteci"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("beklenmeyen imzalama yöntemi: %v", token.Header["alg"])
		}
		return []byte("gizli-anahtar"), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.AuthErrorResponse{
			Message: "Geçersiz yetkilendirme belirteci",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(model.AuthErrorResponse{
			Message: "Geçersiz yetkilendirme belirteci",
		})
	}

	c.Locals("user", token)

	return c.Next()
}
