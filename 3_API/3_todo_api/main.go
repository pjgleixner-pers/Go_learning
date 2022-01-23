package main

import (
	"api/3_todo_api/controller"
	"api/3_todo_api/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
