package main

import (
	"abhiroopsantra.dev/barky/handlers"
	"fmt"
	"net/http"
)

func main() {
	const ServerPort = "8080"
	mux := http.ServeMux{}

	directoriesToBeServed := map[string]string{
		"/app":        "./pages",
		"/app/assets": "./assets",
	}

	for path, directory := range directoriesToBeServed {
		fileServer := http.FileServer(http.Dir(directory))
		mux.Handle(path, http.StripPrefix(path, fileServer))
	}

	mux.HandleFunc("/healthz", handlers.Healthz)

	//fileServer := http.FileServer(http.Dir("./pages"))
	//mux.Handle("/", http.StripPrefix("/", fileServer))

	server := &http.Server{
		Addr:    ":" + ServerPort,
		Handler: &mux,
	}

	fmt.Println("Server is running on port 8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
