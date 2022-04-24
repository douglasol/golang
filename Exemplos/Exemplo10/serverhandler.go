package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "teste"
	p := "conteudo"
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, p)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
