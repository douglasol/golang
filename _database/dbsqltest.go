/*  MS Sql server
 	Onde obter os drivers e instalação?
	https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go

*/
package main

import (
	"context"
	"fmt"
	"log"

	"exemplo8/mydb"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	/* conecta no bd */
	isOk := mydb.MyDbConnect()
	if isOk {
		count, err := readPessoa()
		if err != nil {
			log.Fatal("Error reading Pessoas: ", err.Error())
		}
		fmt.Printf("Read %d row(s) successfully.\n", count)
	} else {
		fmt.Printf("Not Connected! %v", isOk)
	}
}

func readPessoa() (int, error) {
	var count int

	/* Check if database is alive. */
	ctx := context.Background()
	err := mydb.Db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	/* Execute query */
	tsql := "select top 10 pessoaCPF, pessoaNome, pessoaEmail from pessoa;"
	rows, err := mydb.Db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	/* Iterate through the result set */
	for rows.Next() {
		var pessoaCPF, pessoaNome, pessoaEmail string
		/* Get values from row */
		err := rows.Scan(&pessoaCPF, &pessoaNome, &pessoaEmail)
		if err != nil {
			return -1, err
		}
		fmt.Printf("%s,%s,%s\n", pessoaCPF, pessoaNome, pessoaEmail)
		count++
	}
	return count, nil
}
