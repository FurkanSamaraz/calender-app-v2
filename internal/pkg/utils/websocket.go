package ws

import (
	"context"
	"encoding/json"
	"fmt"

	redisrepo "main/internal/pkg/redisrepo"
	api_structures "main/internal/pkg/structures/chat"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/websocket/v2"
)

type Message struct {
	Type string              `json:"type"`
	User string              `json:"user,omitempty"`
	Chat api_structures.Chat `json:"chat,omitempty"`
}

var (
	clients = make(map[string]*api_structures.Client)
)

// define our WebSocket endpoint
func WsHandler(c *websocket.Conn) {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// Add client to clients list
	client := &api_structures.Client{Conn: c, Username: claims["username"].(string)}
	clients[claims["username"].(string)] = client
	fmt.Println("clients", len(clients), clients, c.RemoteAddr())

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	for {
		var msg api_structures.Chat
		err := client.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("error reading json", err)
			break
		}

		fmt.Println("received message", msg)

		// Broadcast message to all connected clients
		for _, c := range clients {
			if c.Username != client.Username {
				err = c.Conn.WriteJSON(msg)
				if err != nil {
					fmt.Println("error writing json", err)
					break
				}
				// Add sender and receiver to each other's contact lists
				redisrepo.AddToContactList(c.Username, client.Username)
				redisrepo.AddToContactList(client.Username, c.Username)
			}
			// Save message to Redis
			chatKey := fmt.Sprintf("chats:%s:%s", msg.From, msg.To)
			chatHistory, err := redisrepo.RedisClient.Get(context.Background(), chatKey).Bytes()
			if err != nil && err != redis.Nil {
				fmt.Println("error fetching chat history", err)
				break
			}

			var messages []api_structures.Chat
			if len(chatHistory) > 0 {
				err = json.Unmarshal(chatHistory, &messages)
				if err != nil {
					fmt.Println("error unmarshaling chat history", err)
					break
				}
			}

			messages = append(messages, msg)

			data, err := json.Marshal(messages)
			if err != nil {
				fmt.Println("error marshaling chat history", err)
				break
			}

			err = redisrepo.RedisClient.Set(context.TODO(), chatKey, data, 0).Err()
			if err != nil {
				fmt.Println("error saving chat history", err)
				break
			}
		}
	}

	fmt.Println("exiting", c.RemoteAddr().String())
	delete(clients, claims["username"].(string))
}
