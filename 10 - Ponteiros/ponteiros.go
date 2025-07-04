package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// PONTEIRO -> REF. DE MEMORIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // * -> desreferenciação

	variavel3 = 150

	fmt.Println(variavel3, *ponteiro)
}
