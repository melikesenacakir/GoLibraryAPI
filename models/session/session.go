package mysess

import "github.com/gofiber/fiber/v2/middleware/session"

type NewSess struct {
	Sess *session.Store
}
