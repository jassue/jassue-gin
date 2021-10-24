package services

import (
    "github.com/dgrijalva/jwt-go"
    "jassue-gin/global"
    "time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

type JwtUser interface {
    GetUid() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
    jwt.StandardClaims
}

const (
    TokenType = "bearer"
)

type TokenOutPut struct {
    AccessToken string `json:"access_token"`
    ExpiresIn int `json:"expires_in"`
    TokenType string `json:"token_type"`
}

// CreateToken 生成 Token
func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
    token = jwt.NewWithClaims(
        jwt.SigningMethodHS256,
        CustomClaims{
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl,
                Id:        user.GetUid(),
                Issuer:    GuardName,
                NotBefore: time.Now().Unix() - 1000,
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