package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

// Nesse caso so conseguimos imprimir o nome
// passado na declaracao
func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// Aqui vamos alterar o valor na memoria
// da variavel que estiver usando essa funcao
// justamente por causa do ponteiro
func (p *pessoa) mudarnome(nome string) {

	partes := strings.Split(nome, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]

}

func main() {

	eu := pessoa{
		nome:      "felipe",
		sobrenome: "tavares",
	}

	fmt.Print(eu.nomeCompleto())
	fmt.Println()

	eu.mudarnome("joao silva")
	fmt.Print(eu.nomeCompleto())
}
