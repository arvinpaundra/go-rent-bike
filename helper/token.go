package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId string, role string) (string, error) {
	claims := jwt.MapClaims{}

	claims["user_id"] = userId
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Cfg.JWTSecret))
}

func ExtractToken(c echo.Context) (interface{}, error) {
	tokenString := c.Request().Header.Get("Authorization")

	formattedTokenString := strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(formattedTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(string)
		role := claims["role"].(string)

		data := map[string]string{
			"user_id": userId,
			"role":    role,
		}

		return data, nil
	}

	return nil, err
}
