package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!(true || false))

	// Unários
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 10
	numero *= 3
	numero /= 3
	numero %= 3
	fmt.Println(numero)

	// Ternário

}
