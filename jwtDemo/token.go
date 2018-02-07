package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type Myclaims struct {
	UserId   int
	Username string
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("hzwy23")
	// Create the Claims
	claims := Myclaims{
		1,
		"wjf",
		jwt.StandardClaims{
			IssuedAt:  int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + 3),
			Issuer:    "wjf",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println("签名后的token信息:", ss)
	time.Sleep(time.Second * 4)
	t, err := jwt.ParseWithClaims(ss, &Myclaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sign method %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println("还原后的token信息claims部分:", t.Claims)
	if c, ok := t.Claims.(*Myclaims); ok {
		fmt.Println(c.UserId)
		fmt.Println(c.Username)
	}
}
