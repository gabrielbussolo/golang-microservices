package service

import (
	"github.com/gabrielbussolo/golang-microservices/mvc/domain"
	"github.com/gabrielbussolo/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
