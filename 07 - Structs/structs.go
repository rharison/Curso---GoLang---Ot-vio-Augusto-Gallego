package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Rharison"
	u.idade = 30

	fmt.Println(u)

	usuario2 := usuario{"Thander", 24, endereco{"Rua", 00}}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
