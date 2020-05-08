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

//MyKey jwt handdler
const (
	SecretKey = "My-temporary-SECRETKey-2017"
)
