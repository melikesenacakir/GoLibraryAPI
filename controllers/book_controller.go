package controllers

import (
	"net/http"
	"rest-api-2/db/dbstore"
	dto "rest-api-2/models"
	mysess "rest-api-2/models/session"
	"rest-api-2/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MyDbandSess struct {
	Query *dbstore.Queries
	Sess  *mysess.NewSess
}

func (query *MyDbandSess) Createbook(c *fiber.Ctx) error {

	dtobook := new(dto.CreateBookDto)
	if err := c.BodyParser(&dtobook); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "invalid request"})
	}

	ctx := c.Context()
	book := dto.CreateDtoToBook(dtobook)
	newquery := &service.MyDb{Query: query.Query}
	err := newquery.CreateBook(book, ctx)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "error adding book"})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
}

func (query *MyDbandSess) Getbook(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	id_int, _ := strconv.Atoi(id)
	newquery := &service.MyDb{Query: query.Query}
	Book, err := newquery.GetBook(id_int, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "error getting book"})
	}
	dtobook := dto.BookToDto(&Book)
	return c.Status(http.StatusOK).JSON(dtobook)
}

func (query *MyDbandSess) Getbooks(c *fiber.Ctx) error {
	ctx := c.Context()
	newquery := &service.MyDb{Query: query.Query}
	Book, err := newquery.GetBooks(ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "error getting books"})
	}
	newbook := dto.BooksToDto(&Book)
	return c.Status(http.StatusOK).JSON(newbook)
}

func (query *MyDbandSess) Deletebook(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	id_int, _ := strconv.Atoi(id)

	newquery := &service.MyDb{Query: query.Query}
	err := newquery.DeleteBook(id_int, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "error deleting book"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been deleted"})
}

func (query *MyDbandSess) Updatebook(c *fiber.Ctx) error {

	dtobook := new(dto.BookDto)
	if err := c.BodyParser(dtobook); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
	}

	ctx := c.Context()
	id, _ := strconv.Atoi(c.Params("id"))
	dtobook.ID = int64(id)
	book := dto.UpdateDtoToBook(dtobook)
	newquery := &service.MyDb{Query: query.Query}
	err := newquery.UpdateBook(*book, ctx)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "error updating book"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been updated"})
}
