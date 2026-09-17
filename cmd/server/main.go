package main

import (
	"fmt"
	"go-curd/internal/handler"
	"net/http"
)

func main() {

	// Simple GET Request...
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Our backend server!")
	})

	// Request using handlers...
	http.HandleFunc("/health", handler.Health)

	fmt.Println("Server is running on :8080")

	//http.ListenAndServe(":8080", nil)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
