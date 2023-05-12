package redisrepo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	api_structure_chat "main/internal/pkg/structures/chat"
	api_structures "main/internal/pkg/structures/employee_requests"

	"github.com/go-redis/redis/v8"
)

func RegisterNewUser(u *api_structures.Employee) error {
	// redis-cli
	// SYNTAX: SET key value
	// SET username password
	// register new username:password key-value pair
	data, err := json.Marshal(u)
	if err != nil {
		log.Println("json", err)
		return err
	}

	err = RedisClient.Set(context.Background(), u.Name, data, 0).Err()
	if err != nil {
		log.Println("error while adding new user", err)
		return err
	}

	return nil
}

func IsUserExist(username string) (bool, error) {
	exists, err := RedisClient.Exists(context.Background(), username).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

func IsUserAuthentic(u *api_structures.Employee) error {
	// redis-cli
	// SYNTAX: GET key
	// GET username

	data, err := RedisClient.Get(context.Background(), u.Name).Bytes()
	if err != nil {
		log.Fatal(err)
	}

	var savedUser api_structures.Employee
	err = json.Unmarshal(data, &savedUser)
	if err != nil {
		log.Fatal(err)
	}

	if savedUser.Password != u.Password {
		log.Fatal("Invalid credentials")
	}

	return nil
}

func FetchChatBetween(username1, username2 string) ([]api_structure_chat.Chat, error) {
	// Construct the Redis key for the chat messages
	//chatKey := fmt.Sprintf("chats:%s:%s", username1, username2)
	// Construct the Redis key for the chat messages
	chatKey := fmt.Sprintf("chats:%s:%s", username1, username2)
	var chatHistory []api_structure_chat.Chat
	// Retrieve the chat messages within the specified time range
	chatData, err := RedisClient.Get(context.Background(), chatKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			// Redis key not found, handle the case accordingly
			fmt.Println("Chat history not found")
		} else {
			// Other error occurred, handle the error
			fmt.Println("Failed to fetch chat history:", err)
		}
	}
	err = json.Unmarshal(chatData, &chatHistory)
	if err != nil {
		log.Fatal(err)
	}
	return chatHistory, nil
}

// FetchContactList of the user. It includes all the messages sent to and received by contact
// It will return a sorted list by last activity with a contact
func FetchContactList(username string) ([]string, error) {
	contactListKey := fmt.Sprintf("contact-list:%s", username)

	// Retrieve the contact list from Redis
	contactList, err := RedisClient.SMembers(context.Background(), contactListKey).Result()
	if err != nil {
		return nil, err
	}

	return contactList, nil
}

func AddToContactList(username, contactUsername string) error {
	// Construct the Redis key for the contact list
	contactListKey := fmt.Sprintf("contact-list:%s", username)

	// Add the contact to the contact list
	err := RedisClient.SAdd(context.Background(), contactListKey, contactUsername).Err()
	if err != nil {
		return err
	}

	return nil
}
