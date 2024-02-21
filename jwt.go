package lab_backend

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func genJWT(object map[string]string) (string, error) {
	claims := &jwt.MapClaims{
		"iss":  "kotoba",
		"exp":  time.Now().Add(time.Duration(Conf().Token.Expiration) * time.Minute).Unix(),
		"data": object,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString([]byte(Conf().Token.PrivateKey))

	if err != nil {
		return "", err
	}

	return res, nil
}

func parseJWT(headerToken string) (*jwt.Token, error) {
	return jwt.Parse(headerToken, func(_token *jwt.Token) (any, error) {
		if _, ok := _token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", _token.Header["alg"])
		}

		return []byte(Conf().Token.PrivateKey), nil
	})
}

func checkJWT(headerToken string) error {
	token, err := parseJWT(headerToken)

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func extractJWT(headerToken string, keys ...string) (map[string]any, error) {
	checkErr := checkJWT(headerToken)

	if checkErr != nil {
		return nil, checkErr
	}

	token, parseErr := parseJWT(headerToken)

	if parseErr != nil {
		return nil, checkErr
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("error casting token.Claims to jwt.MapClaims")
	}

	return claims["data"].(map[string]any), nil
}
