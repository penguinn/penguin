package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	InitError = errors.New("初始化JWT错误")
	config    = JWTConfig{
		secret: "-----oo-0o--o0o--",
	}
)

type JWTConfig struct {
	secret string
}

type JWTComponent struct{}

func (JWTComponent) Init(options ...interface{}) (err error) {
	if len(options) == 0 {
		return InitError
	}
	c, ok := options[0].(*JWTConfig)
	if !ok {
		return InitError
	}
	config = *c
	return nil
}

type Claim struct {
	jwt.StandardClaims
}

func CreateToken(claim Claim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(config.secret))
}

func ParseToken(tokenStr string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*Claim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("没有找到token")
}
