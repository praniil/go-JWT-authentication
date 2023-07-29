// JWT:= json web tokens(can be implemented in any server side programming language (backend))
// we can use golang-jwt package that provides functionality for generating and validating JWTs
// jwt made of 3 parts seperated by "."
// (information like algorithm used to generate the signature).(application specific information like username).(signature)		//JWT
package main

import (
	"log"
	"fmt"
	"net/http"
	"go-jwt-authentication/handlers"
)

func main() {
	//handlers needed
	fmt.Println("serving in :8080 server")
	http.HandleFunc("/signin", handlers.Signin)
	// http.HandleFunc("/welcome", Welcome)
	// http.HandleFunc("/refresh", Refresh)
	// http.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
