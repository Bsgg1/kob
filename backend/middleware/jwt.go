package middleware

import (
	"backend/common"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func GenerateAccessToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Minute * 10).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(common.JWT_secret))

	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func GenerateRefreshToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := token.SignedString([]byte(common.JWT_secret))

	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
func GenerateTokens(userId int64) (*Token, error) {
	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return nil, err
	}
	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return nil, err
	}
	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("access_token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(common.JWT_secret), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID from token"})
			c.Abort()
			return
		}
		userId, ok := claims["user_id"].(int64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID from token"})
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		expTime := time.Unix(int64(claims["exp"].(float64)), 0)
		refreshTime := expTime.Add(-time.Minute * 5) // 在过期前5分钟刷新Token
		if time.Now().After(refreshTime) {
			refreshToken, err := GenerateRefreshToken(userId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
				c.Abort()
				return
			}

			c.Header("New-Token", refreshToken)
		}
		c.Next()
	}
}
