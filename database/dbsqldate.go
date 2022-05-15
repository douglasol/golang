package main

import (
	"database/sql"
	"fmt"
	"modulo/database/mydb"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
)

type Contrato struct {
	Uid uuid.UUID
	Date date
	Company string
}

func main() {
	var contrato Contrato
	contrato.Uid, _ = uuid.NewUUID()
	contrato.Date = time.now()
	contrato.Company = "Empresa"
	contrato.
	if mydb.MyDbConnect() {
		total, _ := contratoINSERT(contrato)
		if total > 0 {
			fmt.Printf("Inserção executada, %v", total)
		}
	} else {
		fmt.Printf("Não conectado!")
	}
}

func contratoINSERT(contrato Contrato) (int64, error) {
	err, ctx := mydb.MyDbContext()
	if err != nil {
		return -1, err
	}

	tsql := `insert into contrato (ContratoGuid, ContratoData, ContratoEmpresa) values (@ContratoGuid);`
	stmt, err := mydb.Db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(
		ctx,
		sql.Named("ContratoGuid", contrato.Uid),
	)
	err = row.Scan()
	if err == nil {
		return -1, err
	}
	return 1, err
}
