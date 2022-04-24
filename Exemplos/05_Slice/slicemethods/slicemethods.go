package slicemethods

/* insere um item na pilha */
func PushItem(s []int, v int) []int {
	return append(s, v)
}

/* insere um item na pilha */
func PopItem(s []int) ([]int, int) {
	v := s[len(s)-1]
	return s[:len(s)-1], v
}

/* insere item a partir de um indice (melhor solução) */
func InsertItemOnIndex(s []int, i int, v int) []int {
	/* caso a posição seja maior que o tamanho do slice */
	if i >= len(s) {
		return append(s, v)
	} else {
		s = append(s[:i+1], s[i:]...)
		s[i] = v
		return s
	}
}

/* remove tratando os indices  (melhor solução) */
func RemoveItemOnIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

/* remove um item da coleção */
func ReplaceItem(s []int, i int, v int) []int {
	if i < len(s) {
		s[i] = v
	}
	return s
}
