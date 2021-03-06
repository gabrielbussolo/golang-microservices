package domain

import (
	"fmt"
	"github.com/gabrielbussolo/golang-microservices/mvc/utils"
	"net/http"
)

var users = map[int64]*User{
	123: &User{123, "Gabriel", "Bussolo", "contact@gabrielbussolo.dev"},
}

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
