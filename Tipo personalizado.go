package main

import "fmt"

// Criamos um apelido para um tipo primitivo
// para podermos criar funcoes com ele
type nome string

// Aqui ja estamos passando uma variavel
// e uma referencia ao mesmo tempo
// diferente de struct
func (n nome) HoM() {

	// Retorna ultima letra como byte
	a := n[len(n)-1]

	// converte byte para string
	// para fazer a operacao logica
	if string(a) == "a" {
		fmt.Println("Mulher")

	} else {
		fmt.Print("Homem")
	}
}

func main() {

	nome.HoM("Jose")

}
