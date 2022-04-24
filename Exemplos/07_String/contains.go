/* basico da linguagem */

package main

func main() {

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
