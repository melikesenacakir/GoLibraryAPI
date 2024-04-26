package config

import "rest-api-2/db"

func DbConfig() *db.DbConf {
	return &db.DbConf{
		Host:     "lib-db",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "librarydb",
	}
}
