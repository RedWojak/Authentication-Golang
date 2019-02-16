package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("password")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprint(w, err.Error())
    }
    
    fmt.Fprintf(w,validToken)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = "6e66b20c-4599-41f7-9665-0bbab5ce4e3e" //UUID
	claims["exp"] = time.Now().Add(time.Second * 3600).Unix()
	claims["orig_iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func main() {
	fmt.Println("JWT test Client")

	handleRequests()

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token")
	}
	fmt.Println(tokenString)

}
