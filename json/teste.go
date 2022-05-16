package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}

type Catlog struct {
	Product_id string `json: "product_id"`
	Quantity   int    `json: "quantity"`
}

func main() {
	file, _ := ioutil.ReadFile("./json/teste.json")

	data := CatlogNodes{}

	err := json.Unmarshal([]byte(file), &data)
	fmt.Printf("%v", err)
	for i := 0; i < len(data.CatlogNodes); i++ {
		fmt.Println("Product Id: ", data.CatlogNodes[i].Product_id)
		fmt.Println("Quantity: ", data.CatlogNodes[i].Quantity)
	}

}
