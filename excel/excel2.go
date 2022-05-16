package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet2")

	xlsx.SetCellValue("Sheet1", "B2", 100)
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetActiveSheet(index)

	if err := xlsx.SaveAs("teste.xlsx"); err != nil {
		fmt.Println(err)
	}
}
