package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado Será retornado")
	fmt.Println("Entrando na funçao para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	// DEFER = Adiar

	defer funcao1()
	funcao2()

	aprovado := alunoEstaAprovado(10, 4)
	fmt.Println("Aprovado: ", aprovado)
}
