package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	connections := make(map[net.Conn]string)

	addNewConnection := make(chan net.Conn)

	postMsg := make(chan string)

	listener, err := net.Listen("tcp", ":31415")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			connection, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			addNewConnection <- connection
		}
	}()

	for {
		select {
		case conn := <-addNewConnection:
			connReader := bufio.NewReader(conn)
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
			go func(conn net.Conn, username string) {
				for {
					message, err := connReader.ReadString('\n')
					if err != nil {
						conn.Close()
						log.Printf("%s disconnected 1", username)
						delete(connections, conn)
						break
					} else {
						log.Printf("Message by %s", username)
						postMsg <- fmt.Sprintf("%s : %s", username, message)
					}
				}
			}(conn, connections[conn])
		case msg := <-postMsg:
			for conn := range connections {
				go func(conn net.Conn) {
					_, err := conn.Write([]byte(msg + "\n"))
					if err != nil {
						conn.Close()
						log.Printf("%s disconnected 2", connections[conn])
						delete(connections, conn)
					}
				}(conn)
			}
		}
	}
}
