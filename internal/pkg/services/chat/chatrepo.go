package service

import (
	"fmt"
	"log"
	"main/internal/pkg/redisrepo"
	api_structure "main/internal/pkg/structures/chat"
)

func (r *ChatService) VerifyContact(username string) (*api_structure.Response, error) {
	// if invalid username and password return error
	// if valid user create new session
	res := &api_structure.Response{Status: true}

	status, err := redisrepo.IsUserExist(username)
	if err != nil {
		log.Fatal(err)
	}
	if !status {
		res.Status = false
		res.Message = "invalid username"
	}

	return res, err
}

func (r *ChatService) ChatHistory(username1, username2 string) (*api_structure.Response, error) {
	// if invalid usernames return error
	// if valid users fetch chats
	res := &api_structure.Response{}

	fmt.Println(username1, username2)
	// // check if user exists
	// if !redisrepo.IsUserExist(username1) || !redisrepo.IsUserExist(username2) {
	// 	res.Message = "incorrect username"
	// 	return res
	// }

	chats, err := redisrepo.FetchChatBetween(username1, username2)
	if err != nil {
		log.Println("error in fetch chat between", err)
		res.Message = "unable to fetch chat history. please try again later."
		return res, err
	}

	res.Status = true
	res.Data = chats
	return res, err
}

func (r *ChatService) ContactList(username string) (*api_structure.Response, error) {
	// if invalid username return error
	// if valid users fetch chats
	res := &api_structure.Response{}

	// // check if user exists
	// if !redisrepo.IsUserExist(username) {
	// 	res.Message = "incorrect username"
	// 	return res
	// }

	contactList, err := redisrepo.FetchContactList(username)
	if err != nil {
		log.Println("error in fetch contact list of username: ", username, err)
		res.Message = "unable to fetch contact list. please try again later."
		return res, err
	}
	fmt.Println(contactList)
	res.Status = true
	res.Data = contactList
	return res, err
}
