package main

import "fmt"

func main() {
	fmt.Println("maps")

	usuario := map[string]string{
		"nome":      "joao",
		"sobrenome": "candido",
	}
	fmt.Println(usuario["nome"], usuario["sobrenome"])
}
