package controller

import (
	"log"

	"main/internal/pkg/redisrepo"
	api_structure_chat "main/internal/pkg/structures/chat"
	api_structures "main/internal/pkg/structures/employee_requests"
)

func Register(u *api_structures.Employee) *api_structure_chat.Response {
	// check if username in userset
	// return error if exist
	// create new user
	// create response for error
	res := &api_structure_chat.Response{Status: true}

	status, err := redisrepo.IsUserExist(u.Name)
	if err != nil {
		log.Fatal(err)
	}
	if status {
		res.Status = false
		res.Message = "username already taken. try something else."
		return res
	}

	err = redisrepo.RegisterNewUser(u)
	if err != nil {
		res.Status = false
		res.Message = "something went wrong while registering the user. please try again after sometime."
		return res
	}

	return res
}

func Login(u *api_structures.Employee) *api_structure_chat.Response {
	// if invalid username and password return error
	// if valid user create new session
	res := &api_structure_chat.Response{Status: true}

	err := redisrepo.IsUserAuthentic(u)
	if err != nil {
		res.Status = false
		res.Message = err.Error()
		return res
	}

	return res
}
