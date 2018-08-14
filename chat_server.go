package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

// Message object
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Color    string `json:"color"`
}

var usernameMap = make(map[string]string)

func handleConnections(so socketio.Socket) {

	so.Join("default-chat-room")
	so.On("joined", func(msg Message) {
		log.Println("A new user connected")
		usernameMap[so.Id()] = msg.Username
		so.Emit("joined", msg)
		so.BroadcastTo("default-chat-room", "joined", msg)
	})
	so.On("message-sent", func(msg Message) {
		log.Println("A new message sent")
		so.Emit("message-sent", msg)
		so.BroadcastTo("default-chat-room", "message-sent", msg)
	})
	so.On("disconnection", func() {
		discUser := usernameMap[so.Id()]
		so.Emit("userDisconnected", discUser)
		so.BroadcastTo("default-chat-room", "userDisconnected", discUser)
	})
}

func main() {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", handleConnections)

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 9000")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
