package jwt

import (
	"fmt"
	"time"

	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/golang-jwt/jwt"
)

// Sign 生成签名
func Sign(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"timestamp": time.Now().UnixMilli(),
	})

	t, err := token.SignedString([]byte(config.GetConfig().JWT.SignKey))
	if err != nil {
		return "", errors.WithMsg(err, "jwt sign token err")
	}

	return t, nil
}

// Parse 解析 token
func Parse(token string) (userID string, valid bool) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetConfig().JWT.SignKey), nil
	})
	if err != nil {
		return "", false
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id, ok := claims["user_id"].(string)
		if ok {
			return id, true
		}
	}

	return "", false
}
