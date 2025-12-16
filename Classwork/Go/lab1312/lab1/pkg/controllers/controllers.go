package controllers

import (
	"encoding/json"
	"lab1312lab1/pkg/models"
	"lab1312lab1/pkg/utils"
	"net/http"
)

func AddUser(writer http.ResponseWriter, request *http.Request) {
	addUser := &models.User{}
	utils.ParseBody(request, addUser)
	user := models.AddUser(addUser)
	res, _ := json.Marshal(user)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
