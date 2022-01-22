package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recived")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recived")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World GO"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
