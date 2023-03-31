package main

import (
	"fmt"
)

func main() {
	/*i := 0
	for i < 5 {
		time.Sleep(time.Second)
		i++
		fmt.Println(i)
	}*/

	nomes := [3]string{"joao", "maria", "jose"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

}
