package main

import (
	"fmt"
	"log"
	"modulo/database/mydb"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	isOk := mydb.MyDbConnect()
	if isOk {
		total, err := pessoaCOUNT()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Total: %v", total)
		}
	} else {
		log.Fatal("desconectado!")
	}
}

func pessoaCOUNT() (int64, error) {
	var total int64
	err, _ := mydb.MyDbContext()
	if err != nil {
		return -1, err
	}
	if err = mydb.Db.QueryRow("select count(*) from pessoa").Scan(&total); err != nil {
		return -1, err
	}
	return total, err
}
