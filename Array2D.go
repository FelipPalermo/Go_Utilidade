package main

// Programa que gera e adiciona x e y
// para dentro de um array / slice 2D infinitamente

import (
	"fmt"
	"math/rand"
	"time"
)

// Funcao para gerar numeros aleatorios
func RN() int {
	number := rand.Intn(25)
	return number
}

func main() {

	// Funcao para random ficar sempre mudando
	rand.Seed(time.Now().UnixNano())

	// coordenadas
	var cord [][2]int

	// variavel de controle
	i := 0
	for {

		if i >= 1 {

			// Caso nao seja a primeira iteracao
			// cord é igual a cord mas sem o primeiro indice
			cord = cord[1:len(cord)]

			i--
		}

		// RA é um vetor 2D que recebe numeros aleatorios
		ra := [2]int{RN(), RN()}

		// Adiciona RA para cord
		cord = append(cord, ra)

		fmt.Println(cord)
		time.Sleep(time.Second * 2)

		i++
		fmt.Printf("\n%d", i)
		fmt.Print("\n")
	}

	fmt.Print(cord)
}
