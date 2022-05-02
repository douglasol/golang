package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mensagem := "<h1>Frontend Simplão</h1>" +
		"<p style='color:red; border-bottom:1px solid black;'>Estou vivo</p>" +
		"<p>Mensagem recebida: <strong>%s</strong></p>" +
		"<script>alert('mensagem legal!')</script>"
	fmt.Fprintf(w, mensagem, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
