/*
package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	// fmt.Println(c2)
	file := excelize.NewFile()
	file.SetCellValue("Sheet1", "A1", "Name")
	file.SetCellValue("Sheet1", "A2", "Dulce")
	file.SetCellValue("Sheet1", "A3", "Mara")

	if err := file.SaveAs("names.xlsx"); err != nil {
		log.Fatal(err)
	}
}

*/
package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)

func main() {
	file, err := excelize.OpenFile("names.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	c1, err := file.GetCellValue("Sheet1", "A2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c1)
	c2, err := file.GetCellValue("Sheet1", "A3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)
}
