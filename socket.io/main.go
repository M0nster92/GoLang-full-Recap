package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello world")

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Emit("change")
		s.Join("bcast")
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		s.Emit("change")
		fmt.Println("event: msg")
		return "recv " + msg
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Serving at localhost:4444...")
	log.Fatal(http.ListenAndServe(":4444", nil))
}
