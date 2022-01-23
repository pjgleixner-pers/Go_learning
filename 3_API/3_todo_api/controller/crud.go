package controller

import (
	"api/3_todo_api/model"
	"api/3_todo_api/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			var data []views.PostRequest
			var err error

			if name != "" {
				data, err = model.ReadByName(name)
			} else {
				data, err = model.ReadAll()
			}

			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} /*
			else if r.Method == http.MethodGet {
				name := r.URL.Query().Get("name")
				data, err := model.ReadByName(name)
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		*/
	}
}
