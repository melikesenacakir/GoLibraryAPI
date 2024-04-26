package controllers

import (
	"net/http"
	dto "rest-api-2/models"
	"rest-api-2/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (query *MyDbandSess) NewBorrow(c *fiber.Ctx) error {
	borrowdto := new(dto.CreateBookborrowDto)
	if err := c.BodyParser(&borrowdto); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "invalid request"})
	}

	ctx := c.Context()
	borrowdto.Status = "borrowed"
	borrowdto.Borrowdate = time.Now()
	borrowdto.Returndate = time.Now().AddDate(0, 0, 10)

	borrow := dto.CreateDtoToBorrow(borrowdto)
	newquery := &service.MyDb{Query: query.Query}
	if err := newquery.NewBorrow(borrow, ctx); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "error borrowing books"})
	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books successfully borrowed. Please return on time!"})

}

func (query *MyDbandSess) OldBorrows(c *fiber.Ctx) error {

	ctx := c.Context()
	status := "returned"
	newquery := &service.MyDb{Query: query.Query}
	history, err := newquery.OldBorrows(status, ctx)
	historydto := dto.BorrowToDto(&history)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "error getting history"})
	}
	return c.Status(http.StatusOK).JSON(historydto)
}

func (query *MyDbandSess) BorrowHistory(c *fiber.Ctx) error {

	ctx := c.Context()
	newquery := &service.MyDb{Query: query.Query}
	history, err := newquery.BorrowHistory(ctx)
	historydto := dto.BorrowHistoryToDto(&history)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "error getting history"})
	}
	return c.Status(http.StatusOK).JSON(historydto)
}

func (query *MyDbandSess) ReturnBorrow(c *fiber.Ctx) error {

	returnbookdto := new(dto.ReturnbookDto)

	ctx := c.Context()
	returnbookdto.Status = "returned"
	id, _ := strconv.Atoi(c.Params("id"))
	returnbookdto.ID = int64(id)

	returnbook := dto.DtoToReturnBorrow(returnbookdto)
	newquery := &service.MyDb{Query: query.Query}
	err := newquery.ReturnBorrow(returnbook, ctx)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "error returning book"})

	}
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book returned"})
}
