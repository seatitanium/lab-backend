package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// 从参数 object 所携带的信息生成一个 JSON Web Token 文本。
func GenerateJWT(object JWTPayload) (string, error) {
	claims := &jwt.MapClaims{
		"iss":  "seati",
		"exp":  time.Now().Add(time.Duration(GlobalConfig.Token.Expiration) * time.Minute).Unix(),
		"data": object,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString([]byte(GlobalConfig.Token.PrivateKey))

	if err != nil {
		return "", err
	}

	return res, nil
}

// 尝试将参数中的 JWT 字符串解析为 *jwt.Token
func ParseJWT(headerToken string) (*jwt.Token, error) {
	return jwt.Parse(headerToken, func(_token *jwt.Token) (any, error) {
		if _, ok := _token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", _token.Header["alg"])
		}

		return []byte(GlobalConfig.Token.PrivateKey), nil
	})
}

// 检查参数中的 JWT 字符串是否有效。按照如下两个方面检查：
//
// 1. 是否可以正确解析。
//
// 2. 可以正确解析时，是否有效（例如是否过期等）。
func CheckJWT(headerToken string) error {
	token, err := ParseJWT(headerToken)

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

// 尝试从参数中的 JWT 字符串中解析出生成时的 Object 内容
func ExtractJWT(headerToken string) (map[string]any, error) {
	checkErr := CheckJWT(headerToken)

	if checkErr != nil {
		return nil, checkErr
	}

	token, parseErr := ParseJWT(headerToken)

	if parseErr != nil {
		return nil, checkErr
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("errHandler casting token.Claims to jwt.MapClaims")
	}

	return claims["data"].(map[string]any), nil
}

func ExtractJWTPayload(headerToken string) *JWTPayload {
	var result JWTPayload

	res, err := ExtractJWT(headerToken)

	if err != nil {
		return nil
	}

	var ok bool

	result.Username, ok = res["username"].(string)
	if !ok {
		return nil
	}

	result.UpdatedAt, ok = res["updated_at"].(int64)
	if !ok {
		return nil
	}

	return &result
}
