package main

import "fmt"

// Funcoes variaticas tem entrada de valores indefinidos de variaveis
// nesse caso ela nos retorna um array / slice de string
func retS(texto ...string) []string {

	return texto
}

func main() {

	texto := retS("ola", "senhor", "barriga")

	// iter e uma funcao variatica que itera sobre um array
	var iter = func(txt ...string) {
		for _, text := range txt {
			fmt.Println(text)
		}
	}

	// os 3 pontos sinalizam que isso se trata de uma variavel slice
	iter(texto...)

}
