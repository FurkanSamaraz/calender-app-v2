package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name" binding:"required"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password" binding:"required"`
	Picture  string    `json:"picture" db:"picture"`
	Lang     string    `json:"lang" db:"lang"`
	Roles    string    `json:"roles" db:"roles"`
	Active   bool      `json:"active" db:"active" default:"true"`
}

type Permissions struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	StoreID uuid.UUID `db:"store_id" json:"store_id"`
}

type Role struct {
	ID     uuid.UUID `db:"id" json:"id"`
	Name   string    `db:"name" json:"name"`
	UserID uuid.UUID `db:"user_id" json:"user_id"`
}

type RolePermissions struct {
	RoleID       uuid.UUID `db:"role_id" json:"role_id"`
	PermissionID uuid.UUID `db:"permission_id" json:"permission_id"`
}

func AuthMiddleware(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Geçersiz yetkilendirme belirteci"})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Geçersiz yetkilendirme belirteci"})
	}

	c.Locals("user", token)
	return c.Next()
}

func AdminRoute(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"].(string) != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Yetersiz izin"})
	}

	return c.JSON(fiber.Map{"message": "Merhaba admin!"})
}

func UserRoute(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"].(string) != "user" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Yetersiz izin"})
	}

	return c.JSON(fiber.Map{"message": "Merhaba kullanıcı!"})
}

func Connection() *gorm.DB {
	dsn := "host=localhost user=postgres password=172754 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
func CheckLogin(c *fiber.Ctx) error {
	db := Connection()
	var users User
	username := c.FormValue("name")
	email := c.FormValue("email")
	db.Table("deep_data.users").Where("name = ? AND email = ?", username, email).First(&users)

	var role Role
	db.Table("deep_data.role").Where("user_id = ?", users.ID).First(&role)

	var rolePermissions RolePermissions
	db.Table("deep_data.role_permissions").Where("role_id = ?", role.ID).First(&rolePermissions)

	// var permisssions Permissions
	// db.Table("deep_data.permissions").Where("id = ?", rolePermissions.PermissionID).First(&permisssions)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": users.Email,
		"role":  role.Name,
	})
	tokenString, err := token.SignedString([]byte("gizli-anahtar"))
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)

	c.Locals("Authorization", tokenString)
	return c.JSON(tokenString)
}

func Login(c *fiber.Ctx) error {
	err := CheckLogin(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Geçersiz yetkilendirme belirteci"})
	}

	return c.JSON(fiber.Map{"message": "Giriş başarılı"})
}

//* main()

// func main() {
// 	app := fiber.New()
// 	app.Get("/login", Login)
// 	app.Use(AuthMiddleware)

// 	app.Get("/admin", AdminRoute)
// 	app.Get("/user", UserRoute)

// 	app.Listen(":3000")
// }
