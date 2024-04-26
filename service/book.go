package service

import (
	"rest-api-2/db/dbstore"

	"github.com/valyala/fasthttp"
)

func (mydb *MyDb) CreateBook(book *dbstore.CreateBookParams, ctx *fasthttp.RequestCtx) error {
	_, err := mydb.Query.CreateBook(ctx, *book)
	return err
}

func (mydb *MyDb) UpdateBook(book dbstore.UpdateBookParams, ctx *fasthttp.RequestCtx) error {
	err := mydb.Query.UpdateBook(ctx, book)
	return err
}

func (mydb *MyDb) GetBook(id int, ctx *fasthttp.RequestCtx) (dbstore.Book, error) {
	Book, err := mydb.Query.GetBook(ctx, int64(id))

	return Book, err
}

func (mydb *MyDb) GetBooks(ctx *fasthttp.RequestCtx) ([]dbstore.Book, error) {
	Book, err := mydb.Query.GetBooks(ctx)

	return Book, err
}
func (mydb *MyDb) DeleteBook(id int, ctx *fasthttp.RequestCtx) error {
	err := mydb.Query.DeleteBook(ctx, int64(id))

	return err
}
