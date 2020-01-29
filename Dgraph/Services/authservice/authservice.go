package authservice

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/dgrijalva/jwt-go/request"
)

//STRUCT DEFINITIONS

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Uid  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

const (
	privKeyPath = "./Services/authservice/keys/app.rsa"
	pubKeyPath  = "./Services/authservice/keys/app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

var verifyBytes, signBytes []byte

func init() {

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("Error reading public key: %v", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}
}

// Login authenticates user when logging in and returns a JWT token
func Login(cred string) string {

	var user UserCredentials

	err := json.Unmarshal([]byte(cred), &user)
	if err != nil {
		fmt.Println("Error in request")
		return ""
	}

	fmt.Println(user.Username, user.Password)

	if strings.ToLower(user.Username) != "test" {
		if user.Password != "test123" {
			fmt.Println("Error logging in")
			fmt.Println("Invalid credentials")
			return ""
		}
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"role": "admin",
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
		"UserInfo": struct {
			Name string
			Role string
		}{user.Username, "Member"}})

	tokenString, err := signer.SignedString(signKey)

	if err != nil {
		fmt.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
	}

	response := Token{tokenString}
	json, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Error parsing respone")
		return ""
	}

	return string(json)
}

//ValidateTokenMiddleware validates the token of the user
// func ValidateTokenMiddleware(token string) {

// if token.Valid {
// 	next(w, r)
// }

//validate token
// token, err := request.ParseFromRequest(r,request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error){
// 	return verifyKey, nil
// })

// if err == nil {

// 	if token.Valid{
// 		next(w, r)
// 	} else { {}
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Token is not valid")
// 	}
// } else {
// 	w.WriteHeader(http.StatusUnauthorized)
// 	fmt.Fprint(w, "Unauthorised access to this resource")
// }

// }
