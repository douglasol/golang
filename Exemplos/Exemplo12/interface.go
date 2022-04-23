package main

import "fmt"

/* -------------------------------------------------------- */
type Artigo struct {
	Titulo string
	Autor  string
}

func (a Artigo) Show() string {
	return fmt.Sprintf("[Artigo]: O artigo %q foi escrito por %s.", a.Titulo, a.Autor)
}

/* -------------------------------------------------------- */
type Livro struct {
	Titulo  string
	Autor   string
	Paginas int
}

func (b Livro) Show() string {
	return fmt.Sprintf("[Livro]: O livro %q foi escrito por %s, e tem %v páginas", b.Titulo, b.Autor, b.Paginas)
}

/* -------------------------------------------------------- */
type Display interface {
	Show() string
}

/* -------------------------------------------------------- */

func main() {
	a := Artigo{
		Titulo: "Understanding Interfaces in Go",
		Autor:  "Sammy Shark",
	}
	b := Livro{
		Titulo:  "Understanding Interfaces in Go",
		Autor:   "Sammy Shark",
		Paginas: 100,
	}

	Print(a)
	Print(b)
}

func Print(s Display) {
	fmt.Println(s.Show())
}
