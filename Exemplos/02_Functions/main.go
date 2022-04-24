/* funções */

package main

//6) global
/*
var nomes []string

func nomesGlobal() {
	nomes = append(nomes, "Zé")
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
}
*/

func main() {

	alomamae()

	//2) maiusculas("alo mamae")

	//3) recebendo um array
	/*
		nomes := nomesGET()
		for _, nome := range nomes {
			fmt.Printf("%v\n", nome)
		}
	*/

	//4) passando um array
	/*
		nomes = nomesToUpperGET(nomes)
		for _, nome := range nomes {
			fmt.Printf("%v\n", nome)
		}
	*/

	//5) recebendo vários retornos de uma vez!!!
	/*
		nomes, nomesupper, nomeslower := nomesListGET()

		for _, nome := range nomes {
			fmt.Printf("%v\n", nome)
		}
		for _, nome := range nomesupper {
			fmt.Printf("%v\n", nome)
		}
		for _, nome := range nomeslower {
			fmt.Printf("%v\n", nome)
		}
	*/

	//6: nomes global (descomente o inicio do programa)
	/*
		nomesGlobal()
		for _, nome := range nomes {
			fmt.Printf("%v\n", nome)
		}
	*/
}

func alomamae() {
	println("alo mamae!")
}

/*
func maiusculas(texto string) {
	println(strings.ToUpper(texto))
}
*/

/*
func nomesGET() []string {
	nomes := []string{}
	nomes = append(nomes, "Zé")
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	return nomes
}
*/

/*
func nomesToUpperGET(nomes []string) []string {
	nomesupper := []string{}
	for _, nome := range nomes {
		nomesupper = append(nomesupper, strings.ToUpper(nome))
	}
	return nomesupper
}
*/

/*
func nomesListGET() ([]string, []string, []string) {
	nomes := []string{}
	nomesupper := []string{}
	nomeslower := []string{}
	// carga
	nomes = append(nomes, "Zé")
	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	// conversao upper
	for _, nome := range nomes {
		nomesupper = append(nomesupper, strings.ToUpper(nome))
	}
	// conversao lower
	for _, nome := range nomes {
		nomeslower = append(nomeslower, strings.ToLower(nome))
	}
	return nomes, nomesupper, nomeslower
}
*/
