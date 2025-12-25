# Controle de Acesso em Go (CLI)

Este é um programa simples escrito em **Go (Golang)** que simula um **controle de acesso** baseado em idade ou convite.

O objetivo é praticar:
 - variáveis
 - tipos básicos
 - impressão no terminal
 - condicionais (`if / else`)
 - operadores lógicos

## 📌 O que o programa faz

O programa:
1. Define um nome, idade e se a pessoa possui convite
2. Mostra essas informações no terminal
3. Verifica se a pessoa:
    - tem **18 anos ou mais**, **OU**
    - possui um **convite**
4. Libera ou nega o acesso com base nessa regra

## codigo fonte

package main

import "fmt"

func main() {

	nome := "João"   // nome de exemplo
	idade := 22     // idade de exemplo
	convite := true // possui convite?

	fmt.Println("=== INFORMAÇÕES ===")
	fmt.Printf("nome: %s\n", nome)
	fmt.Printf("idade: %d\n", idade)
	fmt.Printf("tem convite: %b\n", convite)

	if idade >= 18 || convite == true {
		fmt.Printf("acesso liberado para %s\n", nome)
	} else {
		fmt.Println("acesso negado!")
	}
}
