package main

import (
	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server is starting on ", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
