package main
import "fmt"

func main() {
 var nomes []string{}
 nomes = append(nomes, "Zé")
 nomes = append(nomes, "João")
 nomes = append(nomes, "Maria")
 nomesUpperGET(nomes)
}

func nomesUpperGET ([]string) {
 for _, nome := range nomes {
  fmt.Printf("%v\n", nome)
 }
}
