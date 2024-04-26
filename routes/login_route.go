package routes

import (
	"rest-api-2/controllers"

	"github.com/gofiber/fiber/v2"
)

func (sess *GetSess) LoginRoute(app *fiber.App) {
	myss := controllers.MyDbandSess{Sess: sess.Sess, Query: sess.Query}
	app.Post("/api/v1/login", myss.Login)
}
