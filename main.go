package main

import "fmt"

func main() {
	nome := "João" // nome de exemplo!
	idade := 22 // idade de exemplo!
	convite := true

	fmt.Println("=== INFORMAÇÕES ===")
	fmt.Printf("nome: %s\n", nome)
	fmt.Printf("idade: %d\n", idade)
	fmt.Printf("tem convite: %b\n", convite)

	if  idade >= 18 || convite == true {
		fmt.Printf("acesso liberado para %s\n", nome)
	} else {
		fmt.Println("acesso negado!")
	}
}