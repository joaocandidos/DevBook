package main

import "fmt"

// defer adia a execucao ate antes do retorno da funcao
func alunoAprovado(n1, n2 int) string {
	defer fmt.Println("media calculada com sucesso!!!!!")
	fmt.Println("iniciando .......")

	media := (n1 + n2) / 2

	if media > 6 {
		return "aprovado"
	}
	return "reprovado"
}

func main() {

	fmt.Println(alunoAprovado(4, 6))
}
