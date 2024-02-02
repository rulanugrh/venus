package middleware

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/venus/config"
	"github.com/rulanugrh/venus/internal/entity/dto"
	"github.com/rulanugrh/venus/internal/entity/web"
)

type jwtclaims struct {
	jwt.RegisteredClaims
}

func JWTVerify() fiber.Handler {
	conf := config.GetConfig()
	return jwtware.New(jwtware.Config{
		TokenLookup: "header:Authorization",
		SigningKey: jwtware.SigningKey{Key: []byte(conf.Server.Secret)},
	})
}

func CreateToken(req dto.User) (string, error) {
	conf := config.GetConfig()
	time := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &jwtclaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.Server.Secret))
	if err != nil {
		return "", web.Error{
			Message: err.Error(),
			Code: 500,
		}
	}

	return tokenString, nil
}