package main

import "fmt"

// Struct é uma lista de coisas
// dentro de um objeto
type pessoa struct {
	nome   string
	idade  int
	peso   float32
	altura float32
}

// Aqui estamos referenciando a funcao
// a uma classe / struct
func (p pessoa) imc() float32 {
	return p.peso / (p.altura * 2)
}

func main() {

	// Struct é definido como um dict
	felipe := pessoa{
		nome:   "felipe",
		idade:  22,
		peso:   72.0,
		altura: 178.0,
	}

	fmt.Print(felipe.imc())

}
