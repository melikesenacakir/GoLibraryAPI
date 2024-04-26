package routes

import (
	"rest-api-2/controllers"
	"rest-api-2/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (sess *GetSess) UserRoute(app *fiber.App) {
	myss := middlewares.Mysess{Sess: sess.Sess}
	mydb := &controllers.MyDbandSess{Query: sess.Query}
	users := app.Group("/api/v1/user", myss.AdminAuth)

	users.Post("/create", mydb.Createuser)
	users.Get("/", mydb.Getusers)
	users.Get("/:id", mydb.Getuser)
	users.Delete("/delete/:id", mydb.Deleteuser)
	users.Put("/update/:id", mydb.Updateuser)

}
