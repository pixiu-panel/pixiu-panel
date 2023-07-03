package login

import "github.com/golang-jwt/jwt/v5"

// jwtCustomClaims
// @description: jwtCustomClaims
type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
