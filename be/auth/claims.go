package auth

import 	"github.com/dgrijalva/jwt-go"

type MyClaimsType struct {
	jwt.StandardClaims
	User interface{}	`json:"user"`
}