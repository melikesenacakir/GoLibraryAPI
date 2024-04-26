package routes

import (
	"rest-api-2/controllers"
	"rest-api-2/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (sess *GetSess) BorrowRoute(app *fiber.App) {
	myss := middlewares.Mysess{Sess: sess.Sess}
	mydb := &controllers.MyDbandSess{Query: sess.Query}
	borrow := app.Group("/api/v1/borrows", myss.UserAuth)

	borrow.Post("/new", mydb.NewBorrow)
	borrow.Put("/return/:id", mydb.ReturnBorrow)
	borrow.Get("/history", mydb.BorrowHistory)
	borrow.Get("/oldborrows", mydb.OldBorrows)
}
