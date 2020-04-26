package model

import "github.com/dgrijalva/jwt-go"

type ClaimsJwt struct {
	Username string
	jwt.StandardClaims
}
