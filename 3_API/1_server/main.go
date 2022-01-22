package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recived")
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
