package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Lucas", "Moreira", 45, 189}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UNIP"}
	fmt.Println(e1.nome)
}
