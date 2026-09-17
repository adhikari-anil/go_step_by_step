package main

import (
	"fmt"
	"go-curd/internal/handler"
	"net/http"
)

func main() {

	taskHandler := handler.NewTaskHandlers()

	http.HandleFunc("/", taskHandler.Home)

	http.HandleFunc("/health", taskHandler.Health)

	fmt.Println("Server is running on :8080")

	//http.ListenAndServe(":8080", nil)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
