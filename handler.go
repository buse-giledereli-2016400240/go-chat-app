package clientmicroservice

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	proto "service/proto"

	"github.com/micro/go-micro/client"
)

func SigninCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// call the backend service
	client := proto.NewUserService("go.micro.srv.auth", client.DefaultClient)
	rsp, err := client.Create(context.TODO(), &proto.User{
		Username: request["username"].(string),
		Password: request["password"].(string),
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("Created: %s", rsp.User.Username)

	authResponse, err := client.Auth(context.TODO(), &proto.User{
		Username: request["username"].(string),
		Password: request["password"].(string),
	})

	if err != nil {
		fmt.Println(err)
	}

	response := map[string]interface{}{
		"token": authResponse.Token,
		"valid": authResponse.Valid,
	}

	json.NewEncoder(w).Encode(response)
}

func RestrictedCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// call the backend service
	client := proto.NewUserService("go.micro.srv.auth", client.DefaultClient)
	rsp, err := client.ValidateToken(context.TODO(), &proto.Token{
		Token: request["token"].(string),
		Valid: request["valid"].(bool),
	})
	if err != nil {
		fmt.Println(err)
	}

	if !rsp.Valid {
		json.NewEncoder(w).Encode(rsp.Errors)
	}
	json.NewEncoder(w).Encode(rsp)
}
