package handler

import (
	"context"
	"log"

	proto "chat/clientmicroservice/proto"

	"golang.org/x/crypto/bcrypt"
)

//holds usernames as keys, hashed passwords as values
var passwordMap = make(map[string]string)

//Service is our microservice for authentication
type Service struct{}

//Create is a function for creating a new user
func (ts *Service) Create(ctx context.Context, req *proto.User, res *proto.Response) error {
	if _, ok := passwordMap[req.Username]; ok {
		//if the username already exists in the map an error is created
		res.Error = "Username is already taken!"
		res.Successful = false
		return nil
	}
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Successful = false
		return err
	}
	req.Password = string(hashedPass)
	passwordMap[req.Username] = string(hashedPass)
	res.Successful = true
	return nil
}

//Auth is a function for checking authentication
func (ts *Service) Auth(ctx context.Context, req *proto.User, res *proto.Response) error {
	log.Println("Logging in with:", req.Username)
	if _, ok := passwordMap[req.Username]; ok {
		// Compares our given password against the hashed password stored in the passwordMap
		if err := bcrypt.CompareHashAndPassword([]byte(passwordMap[req.Username]), []byte(req.Password)); err != nil {
			res.Error = "Incorrect password, try again!"
			res.Successful = false
			return nil
		}
		res.Successful = true
		return nil
	}
	//if the username is not on the map it creates an error
	res.Error = "Username does not exist in database!"
	res.Successful = false
	return nil

}
