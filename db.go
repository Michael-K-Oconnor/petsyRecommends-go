package main

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
)

var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost)
var db, err = sql.Open("postgres", dbinfo)
