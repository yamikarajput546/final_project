package userservice

import "github.com/jcromanu/final_project/pkg/entities"

/*
type authenticateResponse struct {
	MessageResponse
	token string
}*/

type createUserResponse struct {
	user    entities.User
	message entities.Message
}

/*
type getUserResponse struct {
	MessageResponse
	entities.User
}

type updateUserResponse struct {
	MessageResponse
}

type deleteUserResponse struct {
	MessageResponse
}
*/