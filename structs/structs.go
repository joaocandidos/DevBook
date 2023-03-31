package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func main() {
	fmt.Println("aquivo struct")

	var u usuario
	u.nome = "joao"
	u.idade = 44
	fmt.Println(u)

	s := usuario{"jose", 345}
	fmt.Println(s)
}
