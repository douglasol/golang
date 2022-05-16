package main

import (
	"fmt"
	"modulo/database/mydb"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
)

type Contrato struct {
	Uid uuid.UUID
}

func main() {

	if mydb.MyDbConnect() {
		contrato, _ := contratoLER()
		fmt.Printf("uid: %s", contrato.Uid.String())
	} else {
		fmt.Printf("Não conectado!")
	}
}

func contratoLER() (Contrato, error) {
	var contrato Contrato
	err, _ := mydb.MyDbContext()
	if err != nil {
		return contrato, nil
	}
	tsql := "select top 1 ContratoGUID from Contrato;"
	rows, err := mydb.Db.Query(tsql)
	if err != nil {
		return contrato, nil
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&contrato.Uid)
		if err != nil {
			return contrato, nil
		}
	}
	return contrato, nil
}
