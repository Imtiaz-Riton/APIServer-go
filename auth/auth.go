package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func JwtBasicAuthentication(hand http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if !ok || UserName != username || Pass != password {
			http.Error(rw, "Access Denied", http.StatusUnauthorized)
			return
		}
		hand.ServeHTTP(rw, req)
	}
}

var SecretKey = []byte(os.Getenv("SecretKey"))
var UserName, Pass string

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, nil
}

func JwtAuthentication(hand http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		tkn := req.Header.Get("Token")

		if len(tkn) == 0 {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Access Denied"))
			return
		}

		token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error occurs")
			}
			return SecretKey, nil
		})
		if err != nil {
			log.Fatal(err)
		}
		if token.Valid {
			hand.ServeHTTP(rw, req)
		}
	}
}

func init()  {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	UserName = os.Getenv("UserName")
	Pass = os.Getenv("PassWord")
}
