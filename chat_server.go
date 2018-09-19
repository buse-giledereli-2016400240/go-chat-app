package main

import (
	proto "chat/clientmicroservice/proto"
	"context"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/micro/go-micro/client"
)

// Message object
type Message struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Message  string `json:"message"`
	Color    string `json:"color"`
}

//map that holds online socket ID's as keys and usernames as values
var usernameMap = make(map[string]string)

func main() {

	//opens a new server
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	//newClient is a client of the auth.proto's UserService
	newClient := proto.NewUserService("go.micro.srv.auth", client.DefaultClient)

	//this function is called when connections are detected by socketio
	server.On("connection", func(so socketio.Socket) {
		//socked is joined to the chat room
		so.Join("default-chat-room")

		//this function is called when messageSent signal is received
		so.On("messageSent", func(msg Message) {
			log.Println("A new message sent")
			//the messageSent signal is sent to every online user
			so.Emit("messageSent", msg)
			so.BroadcastTo("default-chat-room", "messageSent", msg)
		})
		//this function is called when disconnection signal is received
		so.On("disconnection", func() {
			//the disconnected user is deleted from map
			key := so.Id()
			discUser := usernameMap[key]
			delete(usernameMap, key)
			//the userDisconnected signal is sent to every online user
			so.Emit("userDisconnected", discUser)
			so.BroadcastTo("default-chat-room", "userDisconnected", discUser)
		})
		//this function is called when logInRequest signal is received
		so.On("logInRequest", func(msg Message) {
			log.Println("A user requested log in")
			//Auth function is called with given username and password
			rsp, err := newClient.Auth(context.TODO(), &proto.User{
				Username: msg.Username,
				Password: msg.Password,
			})
			if err != nil {
				log.Fatal(err)
			}
			if rsp.Successful {
				log.Println("A user connected")
				//receiveUserList signal is sent with the username of every online user
				for _, value := range usernameMap {
					so.Emit("receiveUserList", value)
				}
				//the new user is added to map
				usernameMap[so.Id()] = msg.Username
				//the joinedSuccessfully signal is sent to the current socket
				so.Emit("joinedSuccessfully", msg)
				//the joined signal is sent to every online user
				so.BroadcastTo("default-chat-room", "joined", msg)
			} else {
				//if there is an error the corresponding error message is sent with the errorJoining signal
				so.Emit("errorJoining", rsp.Error)
			}
		})
		//this function is called when signUpRequest signal is received
		so.On("signUpRequest", func(msg Message) {
			log.Println("A user requested sign up")
			//Create function is called with given username and password
			rsp, err := newClient.Create(context.TODO(), &proto.User{
				Username: msg.Username,
				Password: msg.Password,
			})
			if err != nil {
				log.Fatal(err)
			}
			if rsp.Successful {
				log.Println("A new user connected")
				//receiveUserList signal is sent with the username of every online user
				for _, value := range usernameMap {
					so.Emit("receiveUserList", value)
				}
				//the new user is added to map
				usernameMap[so.Id()] = msg.Username
				//the joinedSuccessfully signal is sent to the current socket
				so.Emit("joinedSuccessfully", msg)
				//the joined signal is sent to every online user
				so.BroadcastTo("default-chat-room", "joined", msg)
			} else {
				//if there is an error the corresponding error message is sent with the errorJoining signal
				so.Emit("errorJoining", rsp.Error)
			}
		})
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Println("Server on port 9000")
	//server is opened on port 9000
	log.Fatal(http.ListenAndServe(":9000", nil))
}
