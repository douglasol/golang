/*
 	Onde obter os drivers e instalação?
	https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go
*/
package mydb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Cursor struct {
	nome  string
	query string
}

/* db connection */
var Db *sql.DB

/* config */
var (
	server   = "DESKTOP-3NV3626"
	port     = 1433
	user     = "sa"
	password = "kjdfnm2"
	database = "VOT44"
)

var Cursores []Cursor

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
	ctx := context.Background()
	err = Db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		isOk = false
	}
	/* conectado */
	return isOk
}
