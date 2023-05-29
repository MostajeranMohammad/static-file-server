package guards

import "github.com/gofiber/fiber/v2"

type JWT interface {
	GetOptionalJWTGuard() fiber.Handler
	GetStrictJWTGuard() fiber.Handler
}
