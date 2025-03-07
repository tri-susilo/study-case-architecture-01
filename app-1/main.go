package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Server is running on port 8090...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
