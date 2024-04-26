package service

import (
	"rest-api-2/db/dbstore"

	"github.com/valyala/fasthttp"
)

func (mydb *MyDb) CreateUser(user *dbstore.CreateUserParams, ctx *fasthttp.RequestCtx) error {
	_, err := mydb.Query.CreateUser(ctx, *user)

	return err
}

func (mydb *MyDb) GetUser(id int, ctx *fasthttp.RequestCtx) (dbstore.User, error) {
	user, err := mydb.Query.GetUser(ctx, int64(id))

	return user, err
}

func (mydb *MyDb) GetUsers(ctx *fasthttp.RequestCtx) ([]dbstore.User, error) {
	user, err := mydb.Query.GetUsers(ctx)

	return user, err
}

func (mydb *MyDb) Deleteuser(id int, ctx *fasthttp.RequestCtx) error {
	err := mydb.Query.DeleteUser(ctx, int64(id))

	return err
}

func (mydb *MyDb) UpdateUser(user *dbstore.UpdateUserParams, ctx *fasthttp.RequestCtx) error {
	err := mydb.Query.UpdateUser(ctx, *user)

	return err
}
