package middleware

import (
	"ApscAdmin/common/config"
	"ApscAdmin/common/constant"
	"ApscAdmin/common/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID  int64  `json:"user_id"`
	Account string `json:"account"`
}

func GenToken(userId int64, account string) (string, error) {
	claims := CustomClaims{
		UserID:  userId,
		Account: account,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ting_ke",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Conf.System.Secret))
}

func ParseToken(tokenStr string) (*CustomClaims, error, string) {
	var errType string
	cc := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, cc, func(token *jwt.Token) (interface{}, error) {
		return config.Conf.System.Secret, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil, ""
	} else {
		if errors.Is(err, jwt.ErrTokenExpired) {
			errType = "expiredErr"
			return nil, err, errType
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			errType = "formatErr"
			return nil, err, errType
		}
		return claims, err, "other"
	}
}

func TokenMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Access-Auth-Token")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, &model.Response{
				Code:    constant.AuthVerifyFailed,
				Status:  "Success",
				Message: "Access failed, invalid token !!! error: token is null",
			})
			return
		} else {
			claims, err, errType := ParseToken(authHeader)
			if err != nil && errType == "formatErr" {
				context.JSON(http.StatusUnauthorized, &model.Response{
					Code:    constant.AuthVerifyFailed,
					Status:  "Success",
					Message: "Access failed, invalid token !!! error: token's format not correct, please check it",
				})
				return
			}
			if err != nil && errType == "expiredErr" {
				context.JSON(http.StatusUnauthorized, &model.Response{
					Code:    constant.AuthVerifyFailed,
					Status:  "Success",
					Message: "Access failed, invalid token !!! error: token's format not correct, please check it",
				})
				return
			}
			if err != nil && errType == "other" {
				context.JSON(http.StatusUnauthorized, &model.Response{
					Code:    constant.AuthVerifyFailed,
					Status:  "Success",
					Message: fmt.Sprintf("Access failed, invalid token !!! error: %s \n", err),
				})
				return
			}
			context.Set("claims", claims)
			context.Next()
		}
	}
}
