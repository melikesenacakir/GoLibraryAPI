package routes

import (
	"rest-api-2/db"
	"rest-api-2/db/dbstore"
	mysess "rest-api-2/models/session"

	"github.com/gofiber/fiber/v2"
)

type GetSess struct {
	Sess  *mysess.NewSess
	Query *dbstore.Queries
}

func SetUp(app *fiber.App, mysess *mysess.NewSess, conf *db.DbConf) {
	temp := dbstore.New(conf.Db)
	dbsess := GetSess{Sess: mysess, Query: temp}

	dbsess.LoginRoute(app)
	dbsess.BookRoute(app)
	dbsess.UserRoute(app)
	dbsess.BorrowRoute(app)
}
