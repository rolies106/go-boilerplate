package models

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId     uint64      `json:"user_id"`
	UserType   string      `json:"user_type"`
	UserDetail interface{} `json:"user_detail"`
	jwt.StandardClaims
}

// Create JWT Token for current buyer
func (this *Token) CreateJWTToken(typeUser string, user interface{}) string {

	// Create new JWT token for the newly registered account
	var id uint64
	switch typeUser {
	case "user_buyers":
		id = user.(*UserBuyers).ID
	}

	tk := &Token{UserId: id, UserType: typeUser, UserDetail: user}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	return tokenString
}
