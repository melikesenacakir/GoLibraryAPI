package service

import (
	"rest-api-2/db/dbstore"

	"github.com/valyala/fasthttp"
)

func (mydb *MyDb) Login(username string, ctx *fasthttp.RequestCtx) (dbstore.LoginRow, error) {
	result, err := mydb.Query.Login(ctx, username)

	return result, err
}
