package models

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	UserID string `json:"user_id"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}

type JwtToken struct {
	Token   string
	Expires string
}

// Key model for JWT authorization
type Key int

//MyKey jwt handdler
const (
	JwtKey    Key = 100000
	SecretKey     = "My-temporary-SECRETKey-2017"
)
