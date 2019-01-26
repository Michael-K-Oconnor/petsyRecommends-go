package main

import (
	"database/sql"
	"fmt"
)

const (
	dbUser     = "MichaelKOconnor"
	dbPassword = "L&e6a5h4c3i2m1"
	dbName     = "ten_mil"
	dbHost     = "mypostgresinstance.cddpmydwbwcw.us-east-1.rds.amazonaws.com"
)

var dbinfo = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost)
var db, err = sql.Open("postgres", dbinfo)
