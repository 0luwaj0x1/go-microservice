package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/microservices/services"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
 userId := 123
 user, _ := services.GetUser(int64(userId))
 //fmt.Println(user)
 jsonValue, _ := json.Marshal(user)
 w.Write(jsonValue)
}