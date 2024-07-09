package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.ServeMux{}

	// Create a new http.Server struct
	server := &http.Server{
		Addr:    ":8080",
		Handler: &mux,
	}

	fmt.Println("Server is running on port 8080")
	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
