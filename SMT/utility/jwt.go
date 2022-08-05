package utility

import (
	responseMessages "SMT/types/strings"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(payload map[string]interface{}) (string, error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	for key, value := range payload {
		claims[key] = value
	}

	jwtToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func GetPayloadFromToken(jwtToken string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(responseMessages.INVALID_JWT_TOKEN)
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(responseMessages.INVALID_JWT_TOKEN)
}
