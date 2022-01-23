package controller

import (
	"api/3_todo_api/model"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//take data
			//save it
			if err := model.CreateTodo(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}
