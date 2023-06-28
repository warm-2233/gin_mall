package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("34s5d67t8y9u0-jjhv") // TODO 写入配置文件

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Authority int    `json:"authority"`

	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)

	claims := Claims{
		ID:        id,
		UserName:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-mall",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok &&
			tokenClaims.Valid { //  && tokenClaims.IssuedAt.Unix() > 0
			return claims, nil
		}
	}
	return nil, err
}

// GenerateEmailToken

type EmailClaims struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`

	jwt.StandardClaims
}

func GenerateEmailToken(id uint, operation uint, email string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)

	claims := EmailClaims{
		ID:            id,
		Email:         email,
		OperationType: operation,
		Password:      password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-mall-email",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseEmailToken(tokenString string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &EmailClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok &&
			tokenClaims.Valid { //  && tokenClaims.IssuedAt.Unix() > 0
			return claims, nil
		}
	}
	return nil, err
}
