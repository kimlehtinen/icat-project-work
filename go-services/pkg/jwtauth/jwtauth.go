package jwtauth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kim3z/icat-project-work/pkg/models"
)

func GenerateToken(user models.User) (models.JwtToken, error) {
	var jwtToken models.JwtToken
	expires := time.Now().Add(time.Minute * 60).Unix()

	userID := user.ID
	isAdmin := user.Role == models.ADMIN_USER
	jwtClaims := models.JwtClaims{
		userID.Hex(),
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: expires,
			Issuer:    "localhost",
		},
	}

	// create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	signedToken, err := token.SignedString([]byte(models.SecretKey))
	if err != nil {
		return models.JwtToken{}, err
	}

	jwtToken.Token = signedToken
	jwtToken.Expires = fmt.Sprintf("%q", time.Unix(expires, 0))

	return jwtToken, nil
}
