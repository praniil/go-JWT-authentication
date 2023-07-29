package main

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var users = map[string]string{
	"user1" : "password1",
	"user2" : "password2",
}

//jwt key to create the signature
var jwtKey = []byte("my_secret_key")

//to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//a struct that will be encoded to a JWT
//jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//creating a signin handler
func Signin(w http.ResponseWriter, r *http.Request){
	var creds Credentials
	//get json body and decode into credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}
