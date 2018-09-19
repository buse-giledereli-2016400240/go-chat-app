package main

import (
	proto "chat/clientmicroservice/proto"
	"chat/handler"
	"fmt"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)
	service.Init()

	// Register handler
	proto.RegisterUserServiceHandler(service.Server(), new(handler.Service))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
