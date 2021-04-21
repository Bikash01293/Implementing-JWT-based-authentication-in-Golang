package main

import (
	"fmt"
	"log"
	
	"time"
	"net/http"
	

	jwt "github.com/dgrijalva/jwt-go"
)
 

var mySigningKey = []byte("mysupersecretphrase")


func homePage(w http.ResponseWriter, r *http.Request){
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}
func GenerateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Bikash"
	claims["exp"] = time.Now().Add(time.Minute*30).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil{
		fmt.Errorf("something went wrong : %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
func main(){
	fmt.Println("my simple client")

	handleRequests()
}	