package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var jwtSecret = []byte("MyVeryVerySecretKey")

// JWTClaims struct for JWT claims
type JWTClaims struct {
	// UserId is the id of user. Need to parse to save it to context
	UserId uuid.UUID `json:"user_id"`
	// Standart JWTClaims
	jwt.StandardClaims
}

// GenerateJWT function for generating JWT token.
func GenerateJWT(userId uuid.UUID) (string, error) {
	claims := &JWTClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(), // 24 Hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT function for validating JWT token.
//
// It returns error if token is not valid or expired.
// If there is no error, it returns JWTClaims.
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

	if !ok {
		return nil, errors.New("invalid JWT claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, err
}
