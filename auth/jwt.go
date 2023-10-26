package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var jwtSecret = []byte("MyVeryVerySecretKey")

type JWTClaims struct {
	UserId uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId uuid.UUID) (string, error) {
	claims := &JWTClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(jwtToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) { return []byte(jwtSecret), nil },
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	fmt.Println(claims.Id)

	if !ok {
		return nil, errors.New("invalid JWT claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, err
}
