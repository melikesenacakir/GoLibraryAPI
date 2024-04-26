package main

import (
	"rest-api-2/config"
	mysess "rest-api-2/models/session"
	"rest-api-2/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New()
	sess := session.New()
	myss := mysess.NewSess{Sess: sess}
	conf := config.DbConfig()
	conf.DbInit()
	routes.SetUp(app, &myss, conf)
	app.Listen(":3000")
}
