package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims is a struct representing custom claims in the JWT token
type CustomClaims struct {
	CustomerInfo struct {
		Name string `json:"name"`
	} `json:"customer_info"`
	jwt.StandardClaims
}

// authHandler is an example authentication handler that produces a token
func GetTokenViaHTTP(serverPort string, verifyKey interface{}) {
	// HTTP Post request to authenticate
	res, err := http.PostForm(fmt.Sprintf("http://localhost:%v/authenticate", serverPort), url.Values{
		"user": {"test"},
		"pass": {"known"},
	})
	fatal(err)

	// Check for unexpected status code
	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code", res.StatusCode)
	}

	// Read the token from the response body
	buf, err := io.ReadAll(res.Body)
	fatal(err)
	res.Body.Close()
	tokenString := strings.TrimSpace(string(buf))

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// Use the public counterpart of the private key to verify
		return verifyKey, nil
	})
	fatal(err)

	// Access claims from the token
	claims := token.Claims.(*CustomClaimsExample)
	fmt.Println(claims.CustomerInfo.Name)
}


// authenticateAndPrintName is a function that performs authentication and prints the customer's name
func UseTokenViaHTTP(serverPort string, verifyKey interface{}) {
	// HTTP Post request to authenticate
	res, err := http.PostForm(fmt.Sprintf("http://localhost:%v/authenticate", serverPort), url.Values{
		"user": {"test"},
		"pass": {"known"},
	})
	fatal(err)

	// Check for unexpected status code
	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code", res.StatusCode)
	}

	// Read the token from the response body
	buf, err := io.ReadAll(res.Body)
	fatal(err)
	res.Body.Close()
	tokenString := strings.TrimSpace(string(buf))

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// Use the public counterpart of the private key to verify
		return verifyKey, nil
	})
	fatal(err)

	// Access claims from the token
	claims := token.Claims.(*CustomClaimsExample)
	fmt.Println(claims.CustomerInfo.Name)
}

// fatal is a helper function to handle errors
func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
    //       // Replace with the actual server port
	// verifyKey := "your_public_key"     // Replace with the actual public key
	// authenticateAndPrintName(serverPort, verifyKey)

	// authHandler(serverPort, verifyKey)
