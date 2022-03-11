package middleware

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/jassue/jassue-gin/app/common/response"
    "github.com/jassue/jassue-gin/app/services"
    "github.com/jassue/jassue-gin/global"
    "strconv"
    "time"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.Request.Header.Get("Authorization")
        if tokenStr == "" {
            response.TokenFail(c)
            c.Abort()
            return
        }
        tokenStr = tokenStr[len(services.TokenType)+1:]

        token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(global.App.Config.Jwt.Secret), nil
        })
        if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
            response.TokenFail(c)
            c.Abort()
            return
        }

        claims := token.Claims.(*services.CustomClaims)
        if claims.Issuer != GuardName {
            response.TokenFail(c)
            c.Abort()
            return
        }

        // token 续签
        if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshGracePeriod {
            lock := global.Lock("refresh_token_lock", global.App.Config.Jwt.JwtBlacklistGracePeriod)
            if lock.Get() {
                err, user := services.JwtService.GetUserInfo(GuardName, claims.Id)
                if err != nil {
                    global.App.Log.Error(err.Error())
                    lock.Release()
                } else {
                    tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
                    c.Header("new-token", tokenData.AccessToken)
                    c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
                    _ = services.JwtService.JoinBlackList(token)
                }
            }
        }

        c.Set("token", token)
        c.Set("id", claims.Id)
    }
}