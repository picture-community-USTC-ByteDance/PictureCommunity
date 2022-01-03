package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SignedKey = "hello world"
)

type UserClaims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

func CreateToken(id int64) (string, error) {
	claims := UserClaims{id, jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
		ExpiresAt: time.Now().Unix() + 8000000000, // 过期时间 7天  配置文件
		Issuer:    "thg",                          // 签名的发行者
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SignedKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParserToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SignedKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("无效的token")
			} else {
				return nil, fmt.Errorf("token不可用")
			}
		}
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token无效")
}
