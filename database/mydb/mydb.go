package mydb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

/* db connection */
var Db *sql.DB

/* config */
var (
	server   = "DESKTOP-3NV3626" //"NCEI-7640"
	port     = 1433
	user     = "sa"
	password = "kjdfnm2"
	database = "GO"
)

/* conecta no database */
func MyDbConnect() bool {
	/* Build connection string */
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	isOk := true

	/* Create connection pool */
	Db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		isOk = false
	}

	err, _ = MyDbContext()
	if err != nil {
		log.Fatal(err.Error())
		isOk = false
	}
	/* conectado */
	return isOk
}

/* trata contexto */
func MyDbContext() (error, context.Context) {
	ctx := context.Background()
	err := Db.PingContext(ctx)
	return err, ctx
}
