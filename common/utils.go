package common

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// VerifyToken parses the signatured token with given secret key and validates the parsed token(claims).
func VerifyToken(tokenString string, claims jwt.Claims, secretKey string) bool {
	keyfunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		} else {
			return []byte(secretKey), nil
		}
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, keyfunc)
	if err != nil {
		return false
	}

	if err := jwtToken.Claims.Valid(); err != nil {
		return false
	}

	return true
}
