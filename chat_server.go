package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//holds a map of connections mainly used to broadcast received messages to every connection
	connections := make(map[net.Conn]string)
	//holds a channel of connection requests
	addNewConnection := make(chan net.Conn)
	//holds a channel of messages to be broadcasted
	postMsg := make(chan string)
	//opens a listener for port 31415
	listener, err := net.Listen("tcp", ":31415")
	if err != nil {
		panic(err)
	}
	//Goroutine running an infinite loop to add new connection requests to addNewConnections channel
	go func() {
		for {
			connection, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			addNewConnection <- connection
		}
	}()

	//Infinite loop to select one of the channels and either add a new client or broadcast a message
	for {
		select {
		//adds a new client and receives messages
		case conn := <-addNewConnection:
			connReader := bufio.NewReader(conn)
			//Prompts the client to choose a username to broadcast the messages in a more organised way
			conn.Write([]byte("Choose username: "))
			username, tooLong, err := connReader.ReadLine()
			if tooLong {
				conn.Write([]byte("Too long, shortened version is set."))
			}
			if err != nil {
				conn.Close()
				break
			}
			log.Printf("Accepted new client, %s", username)
			connections[conn] = string(username)
			//Goroutine running an infinite loop to add new messages to the broadcast channel
			go func(conn net.Conn, username string) {
				for {
					message, err := connReader.ReadString('\n')
					if err != nil {
						//if there is an error connecting, the user disconnected so it is treated accordingly
						conn.Close()
						log.Printf("%s disconnected", username)
						delete(connections, conn)
						break
					} else {
						//message is sent to the postMsg, broadcasting channel
						log.Printf("Message by %s", username)
						postMsg <- fmt.Sprintf("%s : %s", username, message)
					}
				}
			}(conn, connections[conn])
		//Messages on the channel are broadcasted to all clients
		case msg := <-postMsg:
			for conn := range connections {
				//Goroutine writing the message from the channel to all clients
				go func(conn net.Conn) {
					_, err := conn.Write([]byte(msg + "\n"))
					if err != nil {
						conn.Close()
						log.Printf("%s disconnected", connections[conn])
						delete(connections, conn)
					}
				}(conn)
			}
		}
	}
}
