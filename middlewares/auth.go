package middlewares

import (
	"fmt"
	mysess "rest-api-2/models/session"

	"github.com/gofiber/fiber/v2"
)

type Mysess struct {
	Sess *mysess.NewSess
}

func (sess *Mysess) UserAuth(c *fiber.Ctx) error {
	newsess, _ := sess.Sess.Sess.Get(c)
	sessions := newsess.Get("user_id")
	fmt.Println(sessions)
	if sessions != "user" && sessions != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
func (sess *Mysess) AdminAuth(c *fiber.Ctx) error {
	newsess, _ := sess.Sess.Sess.Get(c)
	sessions := newsess.Get("user_id")
	if sessions != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
