package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func handleConnections(so socketio.Socket) {
	log.Println("A new user connected")

	so.Join("chat_room")
	//this function is called when messageSent signal is received
	so.On("messageSent", func(msg string) {
		log.Println("A new message sent")
		//the messageSent signal is sent to every online user
		so.Emit("messageSent", msg)
		so.BroadcastTo("default-chat-room", "messageSent", msg)
	})
}

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", handleConnections)

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
