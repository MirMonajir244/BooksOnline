package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "hallaUlla"

func GenerateToken(username, email string, userid uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"userid":   userid,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VaildateToken(token string) error {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("parse token fail")
	}
	taokenIsvalid := parsedtoken.Valid
	if !taokenIsvalid {
		return errors.New("token is invalid")
	}
	//userID := int64(parsedtoken.Claims.(jwt.MapClaims)["userid"].(float64))
	return nil
}
