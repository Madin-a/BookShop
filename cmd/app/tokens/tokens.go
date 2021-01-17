package tokens

import (
	"fmt"
	"github.com/jwt-go"
)

type MyCustomClaims struct {
	ID int64 `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
	Role string `json:"role"`
	jwt.StandardClaims
}
func CreateToken(id int64, login, password ,role string) string {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := MyCustomClaims{
		id,
		login,
		password,
		role,
		jwt.StandardClaims{
			ExpiresAt: 1500000,
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	fmt.Printf("I am a token = %v\n", ss)
	return ss
}
func ParseToken(tokenString string) *MyCustomClaims {

	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		return nil, nil

	})

	claims := token.Claims.(*MyCustomClaims)

	return claims

}