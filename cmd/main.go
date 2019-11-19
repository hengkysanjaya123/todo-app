package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/hengkysanjaya/todo-app/internal/http"
)

func main() {
	// TODO: Create Container for Controllers

	h := http.Server{
		Addr:    ":8000",
		Handler: server.CompileRouter(),
	}

	fmt.Println("Server running on :8000")
	log.Print(h.ListenAndServe())
}
