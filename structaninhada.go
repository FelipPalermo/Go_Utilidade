package main

import "fmt"

// criamos um struct que vai servir de lista
type item struct {
	nomeItem string
	QTDitem  int
}

// e um struct que vai ser o objeto principal
// e que vai receber item como slice
type inventario struct {
	nomePlayer string
	itens      []item
}

// funcao que itera pelos objetos
// dentro do array
func (i inventario) MostrarItens() {
	for _, item := range i.itens {
		fmt.Print(item)
		fmt.Println()
	}
}
func main() {

	inv := inventario{

		nomePlayer: "Hornui thundercock",
		itens: []item{
			item{"garrafa dagua", 10},
			item{"espada longa", 1},
			item{"pocao de cura", 3},
			item{"band aid", 30},
		},
	}

	inv.MostrarItens()

}
