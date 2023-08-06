package myjwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Role      string `json:"角色"'`
	Platform  string `json:"平台"`
	SecretKey []byte
}

// GenerateJwtToken 生成JWT Token
func (customClaims *CustomClaims) GenerateJwtToken() (string, error) {
	// 构造JWT Token的Payload信息
	new := jwt.RegisteredClaims{
		Subject:   customClaims.Subject,
		ExpiresAt: customClaims.ExpiresAt,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	claims := CustomClaims{
		RegisteredClaims: new,
		Role:             customClaims.Role,
		Platform:         customClaims.Platform,
	}

	// 使用HS256算法，指定秘钥生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(customClaims.SecretKey)
	tokenString, err := token.SignedString(customClaims.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (c *CustomClaims) DecodeJwtToken(tokenString string) error {
	SecretKey := c.SecretKey
	// 解析JWT Token，并使用指定秘钥解码Token
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		// 校验Token中所使用的签名算法HS256是否匹配
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回秘钥，用于校验Token的完整性
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
