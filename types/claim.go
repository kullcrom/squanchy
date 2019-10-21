package types

import (
	"github.com/dgrijalva/jwt-go"
)

//Claims ...
type Claims struct {
	Username string
	jwt.StandardClaims
}
