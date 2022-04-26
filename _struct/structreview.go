package main

import (
	"fmt"
	"math"
)

/* ------------------------------------------ */
type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) extensao() float64 {
	return 2 * math.Pi * c.r
}

/* ------------------------------------------ */
type Rectangle struct {
	a, b float64
}

type Square struct {
	a float64
}

/* ------------------------------------------ */
type Pessoa struct {
	nome string
}

type Robo struct {
	pessoa        Pessoa
	modelo        string
	anoFabricacao int
}

func (p *Pessoa) talk() {
	fmt.Printf("Hello, my name is %v", p.nome)
	return
}

/* ------------------------------------------ */

func main() {

	/* inicialização simples */
	var c Circle
	c2 := Circle{100, 100, 50}
	c3 := Circle{x: 100, y: 100, r: 50}
	fmt.Printf("c=%v, c2=%v, c3=%v\n", c, c2, c3)

	/* inicialização collection*/
	var r []Rectangle
	r1 := []Rectangle{}
	r2 := []Rectangle{{a: 10, b: 100}, {a: 2, b: 3}}
	r3 := []Rectangle{{10, 100}, {2, 3}}
	fmt.Printf("r=%v, r1=%v, r2=%v, r3=%v\n", r, r1, r2, r3)

	/* campos */
	var c4 Circle
	c4.r = 10
	area := math.Pi * math.Pow(c4.r, 2)
	fmt.Printf("area=%v\n", area)

	/* métodos em structs (!muito interessante) */
	var c5 Circle
	c5.r = 10
	area = c5.area()
	extensao := c5.extensao()
	fmt.Printf("area=%v, extensao=%v\n", area, extensao)

	/* embedded (two structs) */
	var pessoa Pessoa
	pessoa.nome = "XPY230"
	var robo Robo
	robo.pessoa = pessoa
	robo.modelo = "HCP11"
	robo.anoFabricacao = 2022
	robo.pessoa.talk()
}
