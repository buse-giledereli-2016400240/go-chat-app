package clientmicroservice

import (
	"context"
	"errors"
	"log"

	proto "service/proto"

	"golang.org/x/crypto/bcrypt"
)

var passwordMap = make(map[string]string)

//Create is a function for creating a new user
func (ts *TokenService) Create(ctx context.Context, req *proto.User, res *proto.Response) error {
	if _, ok := passwordMap[req.Username]; ok {
		// Generates a hashed version of our password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.Password = string(hashedPass)
		passwordMap[req.Username] = string(hashedPass)
		res.User = req
		return nil
	}
	return errors.New("username already taken")

}

//Auth is a function for checking authentication
func (ts *TokenService) Auth(ctx context.Context, req *proto.User, res *proto.Token) error {
	log.Println("Logging in with:", req.Username, req.Password)

	// Compares our given password against the hashed password stored in the passwordMap
	if err := bcrypt.CompareHashAndPassword([]byte(passwordMap[req.Username]), []byte(req.Password)); err != nil {
		return err
	}

	token, err := ts.Encode(req)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

//ValidateToken is a function for validating given token
func (ts *TokenService) ValidateToken(ctx context.Context, req *proto.Token, res *proto.Token) error {

	// Decode token
	claims, err := ts.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println(claims)

	if claims.User.Username == "" {
		return errors.New("invalid user")
	}

	res.Valid = true

	return nil
}
