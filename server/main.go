package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
