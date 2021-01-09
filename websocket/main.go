package main

import "net/http"

func setupRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", ws)
}

func main() {

}
