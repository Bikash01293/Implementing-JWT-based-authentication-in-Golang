package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Sectet Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an errror")
				}
				return mySigningKey, nil
			})

			if err !=nil{
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "NOt Authorized")
		}
	})
}
func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9003", nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}