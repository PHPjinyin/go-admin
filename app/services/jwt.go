package services

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"jin-gin/global"
	"jin-gin/utils"
	"strconv"
	"time"
)

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type jwtServices struct {
}

var JwtServices = new(jwtServices)

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.RegisteredClaims
}

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jwtService *jwtServices) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.App.Config.Jwt.JwtTtl) * time.Second)),
				ID:        user.GetUid(),
				Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))

	tokenData = TokenOutPut{
		tokenStr,
		int(global.App.Config.Jwt.JwtTtl),
		TokenType,
	}
	return
}

func (jwtService *jwtServices) JoinBlackList(token *jwt.Token) (err error) {
	nowTime := time.Now()
	timer := token.Claims.(*CustomClaims).ExpiresAt.Sub(nowTime) * time.Second
	err = global.App.Redis.SetNX(context.Background(), jwtService.getBlackListKey(token.Raw), nowTime, timer).Err()
	return
}

func (jwtService *jwtServices) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + utils.MD5([]byte(tokenStr))
}

// IsInBlacklist token 是否在黑名单中
func (jwtService *jwtServices) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true
}
