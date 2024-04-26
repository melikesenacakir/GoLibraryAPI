package controllers

import (
	"net/http"
	dto "rest-api-2/models"
	"rest-api-2/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func (query *MyDbandSess) Createuser(c *fiber.Ctx) error {
	userdto := new(dto.CreateUserDto)
	if err := c.BodyParser(&userdto); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
	}

	ctx := c.Context()
	pass := HashPassword(userdto.Password)
	userdto.Password = pass
	user := dto.DtoToUser(userdto)

	newquery := &service.MyDb{Query: query.Query}
	err := newquery.CreateUser(user, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "error creating user"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user successfully created"})
}

func (query *MyDbandSess) Getuser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	ctx := c.Context()

	newquery := &service.MyDb{Query: query.Query}
	user, err := newquery.GetUser(id, ctx)
	userdto := dto.UserToDto(&user)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "error getting user"})
	}
	return c.Status(http.StatusOK).JSON(userdto)
}

func (query *MyDbandSess) Getusers(c *fiber.Ctx) error {
	ctx := c.Context()
	newquery := &service.MyDb{Query: query.Query}
	user, err := newquery.GetUsers(ctx)
	userdto := dto.UsersToDto(&user)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "error getting users"})
	}

	return c.Status(http.StatusOK).JSON(userdto)
}

func (query *MyDbandSess) Deleteuser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	ctx := c.Context()

	newquery := &service.MyDb{Query: query.Query}
	err := newquery.Deleteuser(id, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "error deleting user"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user successfully deleted"})
}

func (query *MyDbandSess) Updateuser(c *fiber.Ctx) error {
	userdto := new(dto.UpdateUserDto)

	if err := c.BodyParser(&userdto); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
	}

	id, _ := strconv.Atoi(c.Params("id"))
	user := dto.UpdateDtoToUser(userdto)

	user.ID = int64(id)
	ctx := c.Context()

	newquery := &service.MyDb{Query: query.Query}
	err := newquery.UpdateUser(user, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "error updating user"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user successfully updated"})
}
