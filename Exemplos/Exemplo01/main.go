/* basico da linguagem */
package main

import "fmt"

func main() {

	//input:
	var nome string
	fmt.Println("Forneça seu nome:")
	fmt.Scan(&nome)
	fmt.Printf("Ola %v", nome)

	// concatenacao:
	/*
		var nome string
		var sobrenome string
		var nomeinteiro = nome + " " + sobrenome
		fmt.Println(nomeinteiro)
	*/

	// array:
	/*
		var booking [50]string
		booking[0] = "primeiro"
		booking[1] = "segundo"
		booking[2] = "terceiro"
		fmt.Printf("Tamanho = %v\n", len(booking))
		fmt.Printf("Tipo = %T\n", booking)
	*/

	// slices:
	/*
		var booking []string
		booking := []string{}
		booking = append(booking, "primeiro")
		booking = append(booking, "segundo")
		booking = append(booking, "terceiro")
		fmt.Printf("Valores = %v\n", booking)
	*/

	// loops:
	/*
		nomecompleto := []string{}
		primeironome := []string{}
		// entrada
		for {
			var inome string
			var isobrenome string
			fmt.Println("Forneça seu nome (. encerrra):")
			fmt.Scan(&inome)

			continuar := inome == "."
			if continuar {
				break
			} else {
				fmt.Println("Forneça seu sobrenome:")
				fmt.Scan(&isobrenome)
				nomecompleto = append(nomecompleto, inome+" "+isobrenome)
			}
		}
		for _, nome := range nomecompleto {
			var nomes = strings.Fields(nome)
			primeironome = append(primeironome, nomes[0])
		}
		for i, nome := range primeironome {
			fmt.Printf("Nome[%v]=%v\n", i, nome)
		}
	*/

	// validação de input
	/*
		for {
			var inome string
			var isok bool

			fmt.Print("Nome (. encerrra):")
			fmt.Scan(&inome)

			// strings vazias
			match1, _ := regexp.MatchString("^$", inome)
			fmt.Printf("vazia: %v\n", match1)

			// tamanho da string
			isok = len(inome) >= 2
			fmt.Printf("tamanho: %v\n", isok)

			// somente letras e espaço
			match2, _ := regexp.MatchString("/^[a-zA-Z]*$/", inome)
			fmt.Printf("somente letras e espaço: %v\n", match2)
		}
	*/

	// Contains
	/*
		var icidade string
		cidades := "SãoPaulo, Londres, Amesterdan"

		fmt.Print("Cidade:")
		fmt.Scan(&icidade)

		isIn := strings.Contains(strings.ToUpper(cidades), strings.ToUpper(icidade))
		fmt.Printf("Achou: %v", isIn)
	*/

	// Split
	/*
		listacidades := "SãoPaulo;Londres;Amesterdan"
		cidades := []string{}
		cidades = strings.Split(listacidades, ";")
		for _, cidade := range cidades {
			fmt.Printf(": %v\n", cidade)
		}
	*/

	// strings.Map

	/*rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	mapa := func(r rune) rune {
		switch r {
		case 'a', 'e', 'O':
			return '@'
		default:
			return r
		}
	}
	fmt.Println(strings.Map(mapa, "Ola mamae!..."))
	*/
}
