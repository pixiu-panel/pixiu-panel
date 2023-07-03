package cache

import "github.com/golang-jwt/jwt/v5"

// JwtCustomClaims
// @description: jwtCustomClaims
type JwtCustomClaims struct {
	Id         string `json:"id"`         // 用户Id
	Username   string `json:"username"`   // 账号
	IsVerified bool   `json:"isVerified"` // 是否验证邮箱
	jwt.RegisteredClaims
}
