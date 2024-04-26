package controllers

import (
	"net/http"
	dto "rest-api-2/models"
	"rest-api-2/service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (mysess *MyDbandSess) Login(c *fiber.Ctx) error {
	userdto := new(dto.LoginDto)
	if err := c.BodyParser(&userdto); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "invalid request"})
	}
	ctx := c.Context()
	user := dto.DtoToLogin(userdto)

	dblogin := service.MyDb{Query: mysess.Query}

	result, err := dblogin.Login(user.Username, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "login failed"})
	}
	login := dto.LoginToDto(&result)
	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user.Password))
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "login failed"})
	}

	sess, _ := mysess.Sess.Sess.Get(c)
	sess.Set("user_id", result.Role)
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "logged in successfully"})
}
