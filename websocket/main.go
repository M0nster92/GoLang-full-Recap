package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte("Hello from socket")); err != nil {
			log.Fatal(err)
			return
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Client connected .....")

	reader(ws)
}

func main() {
	fmt.Println("Web socket")
	//fmt.Println(time.Now().Unix())
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
