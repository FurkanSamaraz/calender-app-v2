package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Room struct {
	ID       string                   `json:"id"`
	Clients  map[*websocket.Conn]bool `json:"-"`
	Messages chan []byte              `json:"-"`
	mu       sync.Mutex
}

type WebsocketMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

var (
	rooms map[string]*Room
	mu    sync.Mutex
)

func NewRoom() *Room {
	return &Room{
		ID:       generateID(),
		Clients:  make(map[*websocket.Conn]bool),
		Messages: make(chan []byte),
	}
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func generateRoomID() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (r *Room) AddClient(client *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Clients[client] = true
}

func (r *Room) RemoveClient(client *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Clients, client)
}

func (r *Room) Broadcast(message []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for client := range r.Clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Websocket write error:", err)
			r.RemoveClient(client)
		}
	}
}

func StreamWebsocket(c *websocket.Conn) {
	roomID := strings.TrimPrefix(c.Params("roomID"), "/")
	room, ok := rooms[roomID]
	if !ok {
		log.Printf("Room not found: %s", roomID)
		c.Close()
		return
	}

	room.AddClient(c)

	// Send room ID to the client
	err := c.WriteMessage(websocket.TextMessage, []byte(roomID))
	if err != nil {
		log.Println("Websocket write error:", err)
	}

	// Send existing messages to the new client
	for message := range room.Messages {
		err := c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Websocket write error:", err)
			break
		}
	}

	// Handle incoming messages from clients
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Websocket read error:", err)
			break
		}

		log.Println("Received message:", string(message)) // Add this line to log the received message

		var msg WebsocketMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Failed to parse message:", err)
			continue
		}

		switch msg.Event {
		case "offer", "answer", "candidate":
			// Broadcast the message to all clients in the room
			room.Broadcast(message)

		default:
			log.Println("Invalid message event:", msg.Event)
		}
	}

	room.RemoveClient(c)
}

func createRoom(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	room := NewRoom()
	roomID := generateRoomID()

	rooms[roomID] = room

	go func() {
		for {
			select {
			case message := <-room.Messages:
				room.Broadcast(message)
			}
		}
	}()

	return c.SendString(roomID)

}
func main() {
	rooms = make(map[string]*Room)
	app := fiber.New()

	app.Get("/room/create", createRoom)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})
	app.Get("/room/:roomID", func(c *fiber.Ctx) error {
		return c.SendFile("./public/room.html")
	})
	app.Get("/room/:roomID/websocket", websocket.New(StreamWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))

	app.Static("/", "./public")

	log.Fatal(app.Listen(":8000"))
}
