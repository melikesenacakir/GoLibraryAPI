package routes

import (
	"rest-api-2/controllers"
	"rest-api-2/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (sess *GetSess) BookRoute(app *fiber.App) {
	myss := middlewares.Mysess{Sess: sess.Sess}

	mydb := &controllers.MyDbandSess{Query: sess.Query}
	app.Get("/api/v1/book/:id", mydb.Getbook)
	app.Get("/api/v1/book/", mydb.Getbooks)
	books := app.Group("/api/v1/book", myss.AdminAuth)

	books.Post("/create", mydb.Createbook)
	books.Delete("/delete/:id", mydb.Deletebook)
	books.Put("/update/:id", mydb.Updatebook)
}
