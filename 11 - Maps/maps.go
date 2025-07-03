package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Rharison",
		"sobrenome": "Abreu",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"endereco": {
			"logradouro": "Rua X",
		},
		"curso": {
			"nome": "Sistemas para internet",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}

	fmt.Println(usuario2)
}
