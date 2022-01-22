package main

import (
	"api/3_todo_api/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
