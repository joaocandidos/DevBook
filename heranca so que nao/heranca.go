package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    int
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"jose", "oliveira", 44, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia de computacao", "univesp"}
	fmt.Println(e1.altura)
}
