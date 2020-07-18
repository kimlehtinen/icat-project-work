package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kim3z/icat-project-work/pkg/models"
)

/*
	The code in this file is taken from https://github.com/thiagozs/api-golang-jwt/blob/master/api/middlewares/token.go

	The MIT License (MIT)
	Copyright (c) 2017 THIAGO ZILLI SARMENTO
	Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
	The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// Auth is a middleware for authenticating users
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var tokenStr string

		// Get token from query params
		tokenStr = r.URL.Query().Get("jwt")

		// Get token from authorization header
		if tokenStr == "" {
			bearer := r.Header.Get("Authorization")
			if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
				tokenStr = bearer[7:]
			}
		}

		// Get token from cookie
		if tokenStr == "" {
			cookie, err := r.Cookie("jwt")
			if err == nil {
				tokenStr = cookie.Value
			}
		}

		//if token not recovered on the last 3 steps
		if tokenStr == "" {
			log.Printf("[RequireTokenAuthentication] Not have Token to validate")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(models.SecretKey), nil
		})

		if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
			log.Printf("[RequireTokenAuthentication] Token valid! Go forward")
			ctx := context.WithValue(r.Context(), models.JwtKey, *claims)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else if ve, ok := err.(*jwt.ValidationError); ok {

			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Printf("[RequireTokenAuthentication] Thats not even token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				log.Printf("[RequireTokenAuthentication] Token is expired or not active")
			} else {
				log.Printf("[RequireTokenAuthentication] Couldn't handle token: %q", err)
			}

		} else {
			log.Printf("[RequireTokenAuthentication] Couldn't handle token: %q", err)
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
