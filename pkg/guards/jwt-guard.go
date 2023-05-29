package guards

import (
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type jwtGuard struct {
	signingKey string
}

func NewJWTGuard(signingKey string) *jwtGuard {
	return &jwtGuard{
		signingKey,
	}
}

func (jg jwtGuard) GetStrictJWTGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jg.signingKey)},
	})
}

func (jg jwtGuard) GetOptionalJWTGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jg.signingKey)},
		Filter: func(c *fiber.Ctx) bool {
			tokenString := c.GetReqHeaders()["Authorization"]
			return len(strings.Split(tokenString, " ")) < 2
		},
	})
}
