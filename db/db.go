package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DbConf struct {
	Host, User, Password, Port, Dbname string
	Db                                 *sql.DB
}

func (conf *DbConf) DbInit() {
	var err error
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.User, conf.Password, conf.Dbname)
	conf.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection failed")
	}
}
