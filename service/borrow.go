package service

import (
	"rest-api-2/db/dbstore"

	"github.com/valyala/fasthttp"
)

func (mydb *MyDb) NewBorrow(borrow *dbstore.BookBorrowParams, ctx *fasthttp.RequestCtx) error {
	_, err := mydb.Query.BookBorrow(ctx, *borrow)
	return err
}

func (mydb *MyDb) OldBorrows(status string, ctx *fasthttp.RequestCtx) ([]dbstore.GetBorrowsRow, error) {
	history, err := mydb.Query.GetBorrows(ctx, status)
	return history, err
}

func (mydb *MyDb) BorrowHistory(ctx *fasthttp.RequestCtx) ([]dbstore.GetBorrowHistoryRow, error) {
	history, err := mydb.Query.GetBorrowHistory(ctx)
	return history, err
}

func (mydb *MyDb) ReturnBorrow(returnbook *dbstore.ReturnbookParams, ctx *fasthttp.RequestCtx) error {
	err := mydb.Query.Returnbook(ctx, *returnbook)
	return err
}
