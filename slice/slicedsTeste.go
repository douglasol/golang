package main

import (
	"fmt"
	"modulo/slice/sliceds"
)

func main() {

	/* array */
	a := [4]int{10, 20, 30, 40}
	fmt.Printf("array:%v\n", a)

	/* slice */
	s := []int{10, 20, 30, 40}
	fmt.Printf("slice:%v\n", s)

	/* append in slice */
	s = append(s, 50)
	s = append(s, 60)
	s = append(s, 70)
	s = append(s, 80)
	s = append(s, 90)
	s = append(s, 100)
	fmt.Printf("append:%v\n", s)

	/* push a item in stack */
	fmt.Printf("remove, insert, replace\n")

	/* remove item da coleção */
	i := 2
	fmt.Printf("remove: s=%v, i=%v, s[i]=%v\n", s, i, s[i])
	s = sliceds.RemoveItemOnIndex(s, i)
	fmt.Printf("remove(index):%v\n", s)

	/* insere um novo na posição */
	v := 200 // valor a inserir
	fmt.Printf("insert: s=%v, i=%v, s[i]=%v, v=%v\n", s, i, s[i], v)
	s = sliceds.InsertItemOnIndex(s, i, v)
	fmt.Printf("insert:%v\n", s)

	/* replace */
	i = 2   // posicao a trocar
	v = 300 // valor a trocar
	fmt.Printf("replace: s=%v, i=%v, s[i]=%v, v=%v\n", s, i, s[i], v)
	s = sliceds.ReplaceItem(s, i, v)
	fmt.Printf("replace:%v\n", s)

	/* push and pop a item in stack */
	fmt.Printf("pilha\n")

	//Push(Object)	Inserts an object at the top of the Stack.
	v = 400
	fmt.Printf("push:stack=%v, v=%v\n", s, v)
	s = sliceds.StackPushItem(s, v)
	fmt.Printf("push:stack=%v\n", s)

	//Pop()			Removes and returns the object at the top of the Stack.
	fmt.Printf("pop:stack=%v\n", s)
	s, v = sliceds.StackPopItem(s)
	fmt.Printf("pop:stack=%v, v=%v\n", s, v)

	fmt.Printf("pop:stack=%v\n", s)
	s, v = sliceds.StackPopItem(s)
	fmt.Printf("pop:stack=%v, v=%v\n", s, v)

	fmt.Printf("pop:stack=%v\n", s)
	s, v = sliceds.StackPopItem(s)
	fmt.Printf("pop:stack=%v, v=%v\n", s, v)
}
