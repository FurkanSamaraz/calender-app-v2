package structures

import "github.com/gofiber/websocket/v2"

type StatusMessage struct {
	Message string `json:"message"`
}

func (p *StatusMessage) TableName() string {
	return "calendar.users"
}

type Chat struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}
type Client struct {
	Conn     *websocket.Conn
	Username string `json:"username"`
}
type ContactList struct {
	Username     string `json:"username"`
	LastActivity int64  `json:"last_activity"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *User) TableName() string {
	return "calendar.users"
}
