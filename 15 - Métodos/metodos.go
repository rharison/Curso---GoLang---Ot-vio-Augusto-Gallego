package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u *usuario) salvar() {
	fmt.Printf("Sanvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u *usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()
	fmt.Println("Maior de idade: ", usuario1.maiorDeIdade())

	usuario2 := usuario{"Usuário 2", 17}
	usuario2.salvar()
	usuario2.fazerAniversario()
	fmt.Println("Maior de idade: ", usuario2.maiorDeIdade())
	fmt.Println(usuario2)
}
