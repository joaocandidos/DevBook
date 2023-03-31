package main

import "fmt"

func main() {
	s := []int{4, 5, 6}
	fmt.Println(s)
	r := append(s, 454, 0001)
	fmt.Println(r)
	// [0:2] pegando so da posicao 0 ate a posicao 1
	w := r[0:2]
	fmt.Println(w)
	q := [4]int{0, 1, 2, 3}
	t := q[1:2]
	fmt.Println(t)

	slice := make([]float32, 5, 10)
	//funcao len mostra o tamanho
	fmt.Println("tamanho: ", len(slice))
	//funcao cap mostra a capacidade
	fmt.Println("capacidade: ", cap(slice))
}
