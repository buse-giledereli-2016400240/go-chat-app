package clientmicroservice

import (
	proto "service/proto"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// A randomly created secure key string
var key = []byte("608bf6601500922e3c3cb0f7c065f012")

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *proto.User
	jwt.StandardClaims
}

// Authable is an interface for encodable and decodable users
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *proto.User) (string, error)
}

// TokenService is our main service
type TokenService struct{}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *proto.User) (string, error) {
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "go.micro.srv.auth",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	// Parse the token
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	}
	return nil, err

}
