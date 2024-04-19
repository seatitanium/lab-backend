package utils

import (
	"fmt"
	"seatimc/backend/errHandler"
	"time"

	"github.com/golang-jwt/jwt"
)

// 生成一个 JSON Web Token 文本，Claim 的 data 部分为参数 object
func GenerateJWT(object JWTPayload) (string, *errHandler.CustomErr) {
	claims := &jwt.MapClaims{
		"iss":  "seati",
		"exp":  time.Now().Add(time.Duration(GlobalConfig.Token.Expiration) * time.Minute).Unix(),
		"data": object,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString([]byte(GlobalConfig.Token.PrivateKey))

	if err != nil {
		return "", errHandler.ServerError(err)
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
//  1. 是否可以解析，否则返回 BadToken
//  2. 可以正确解析时，是否有效（例如是否过期等），否则返回 InvalidToken
func CheckJWT(headerToken string) *errHandler.CustomErr {
	token, err := ParseJWT(headerToken)

	if err != nil {
		// 无法解析 token 内容
		return errHandler.BadToken()
	}

	if !token.Valid {
		// token 无效
		return errHandler.InvalidToken()
	}

	return nil
}

// 尝试从 JWT Token 字符串中解析出 payload，返回一个 map[string]any
func ExtractJWT(headerToken string) (map[string]any, *errHandler.CustomErr) {
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
		return nil, errHandler.ServerError(fmt.Errorf("errHandler casting token.Claims to jwt.MapClaims"))
	}

	return claims["data"].(map[string]any), nil
}

// 尝试从 JWT Token 字符串中解析出 payload，并尝试将其转换为 *JWTPayload。如果转换失败，返回 nil
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

	result.UpdatedAt, ok = res["updatedAt"].(int64)
	if !ok {
		return nil
	}

	return &result
}
