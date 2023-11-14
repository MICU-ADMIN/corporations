```go

package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
// JwtWrapper is a struct that holds the secret key, issuer and expiration time for a JWT token
type JwtWrapper struct {
 SecretKey       string // SecretKey is the secret key used for signing the JWT token
 Issuer          string // Issuer is the issuer of the JWT token
 ExpirationMinutes int64 // ExpirationMinutes is the number of minutes the JWT token will be valid for
 ExpirationHours int64 // ExpirationHours is the expiration time of the JWT token in hours
}

// JwtClaim adds email as a claim to the token
// JwtClaim is a struct that holds the Email of the user, as well as the StandardClaims
type JwtClaim struct {
 Email    string     
 jwt.StandardClaims 
}

// GenerateToken generates a jwt token
// GenerateToken takes an email as an argument and returns a signed JWT token and an error
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
 claims := &JwtClaim{
  Email: email,
  StandardClaims: jwt.StandardClaims{
   ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
   Issuer:    j.Issuer,
  },
 }
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 signedToken, err = token.SignedString([]byte(j.SecretKey))
 if err != nil {
  return
 }
 return
}

// RefreshToken generates a refresh jwt token
// RefreshToken takes an email as an argument and returns a signed JWT token and an error
func (j *JwtWrapper) RefreshToken(email string) (signedtoken string, err error) {
 claims := &JwtClaim{
  Email: email,
  StandardClaims: jwt.StandardClaims{
   ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
   Issuer:    j.Issuer,
  },
 }
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 signedtoken, err = token.SignedString([]byte(j.SecretKey))
 if err != nil {
  return
 }
 return
}

//ValidateToken validates the jwt token
// ValidateToken takes a signed JWT token as an argument and returns the JwtClaim and an error
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
 token, err := jwt.ParseWithClaims(
  signedToken,
  &JwtClaim{},
  func(token *jwt.Token) (interface{}, error) {
   return []byte(j.SecretKey), nil
  },
 )
 if err != nil {
  return
 }
 claims, ok := token.Claims.(*JwtClaim)
 if !ok {
  err = errors.New("Couldn't parse claims")
  return
 }
 if claims.ExpiresAt < time.Now().Local().Unix() {
  err = errors.New("JWT is expired")
  return
 }
 return
}

```

```mermaid

Here is a Mermaid Markdown overview for the Go file you provided:
```mermaid
graph LR
    A[JwtWrapper] --> B[SecretKey]
    B --> C[Issuer]
    B --> D[ExpirationMinutes]
    B --> E[ExpirationHours]
    A --> F[GenerateToken]
    F --> G[Claims]
    G --> H[SignedToken]
    A --> I[RefreshToken]
    I --> J[Claims]
    J --> K[SignedToken]
    K --> L[ValidateToken]
    L --> M[Claims]
    M --> N[Error]
```
In this overview, the `JwtWrapper` struct is at the top of the graph, representing the root of the file. It has several dependencies:
* `SecretKey`: a string that holds the secret key used for signing the JWT token.
* `Issuer`: a string that holds the issuer of the JWT token.
* `ExpirationMinutes`: an integer that holds the number of minutes the JWT token

```
