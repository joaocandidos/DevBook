package main

import "fmt"

func main() {
	fmt.Println("ponteiros")
	var v int = 10
	var s int = v
	v++
	fmt.Println(v, s)
	//ponteiro é um endereco de memoria
	var va1 int
	var va2 *int
	fmt.Println(&va1, &va2)
}
