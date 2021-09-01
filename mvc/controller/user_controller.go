package controller

import (
	"encoding/json"
	"github.com/gabrielbussolo/golang-microservices/mvc/service"
	"github.com/gabrielbussolo/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		marshal, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.Status)
		response.Write(marshal)
		return
	}

	user, apiError := service.GetUser(userId)
	if apiError != nil {
		marshal, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.Status)
		response.Write(marshal)
		return
	}

	jsonValue, _ := json.Marshal(user)

	response.Write(jsonValue)
}
