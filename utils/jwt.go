package utils

import (
	"encoding/json"
	"fmt"
	"seatimc/backend/errHandler"
	"time"

	"github.com/golang-jwt/jwt"
)

// 生成一个 JSON Web Token 文本，Claim 的 data 部分为参数 object
func GenerateJWT(object JWTPayload) (string, *errHandler.CustomErr) {

	claims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "seati",
			ExpiresAt: time.Now().Add(time.Duration(GlobalConfig.Token.Expiration) * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Payload: object,
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

// 提取出 JWT Token 字符串中的 claims 部分，返回一个 jwt.MapClaims（相当于 map[string]any）
func GetClaimsFromToken(headerToken string) (jwt.MapClaims, *errHandler.CustomErr) {
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
		return nil, errHandler.ServerError(fmt.Errorf("cannot convert claims"))
	}

	return claims, nil
}

func GetPayloadFromToken(headerToken string) (*JWTPayload, *errHandler.CustomErr) {
	claims, customErr := GetClaimsFromToken(headerToken)

	payloadData := claims["payload"].(map[string]any)
	payload := JWTPayload{}

	marshaled, err := json.Marshal(payloadData)

	if err != nil {
		return nil, errHandler.ServerError(err)
	}

	err = json.Unmarshal(marshaled, &payload)

	if err != nil {
		return nil, errHandler.ServerError(err)
	}

	if customErr != nil {
		return nil, customErr
	}

	return &payload, nil
}
