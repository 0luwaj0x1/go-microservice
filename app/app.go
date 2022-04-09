package app

import (
	"net/http"

	"example.com/microservices/controllers"
)


func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}