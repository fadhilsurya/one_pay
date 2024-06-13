package helper

import (
	"fmt"
	"one_pay/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTHelper struct {
	jwtKey []byte
}

func NewJWTHelper() *JWTHelper {
	return &JWTHelper{
		jwtKey: []byte(config.AppConfig.JWT.Secret),
	}
}

type Claims struct {
	UserCode string `json:"user_code"`
	jwt.StandardClaims
}

func (j *JWTHelper) GenerateJWT(userCode string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserCode: userCode,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTHelper) VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
