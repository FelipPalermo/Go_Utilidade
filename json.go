package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Criando struct com estrutura para json
// quando usamos `json:""`
// Podemos indicar como esse parametro vai convertido para json
type jogador struct {
	Id     int    `json:"ID"`
	Nome   string `json:"Nome"`
	Classe string `json:"Classe"`
	Level  int    `json:"Level"`
}

func CriarJogador() {

	// input é a variavel responsavel por fazer a entrada
	// do que o usuario escreve
	input := bufio.NewReader(os.Stdin)
	Nome := ""
	Classe := ""
	arquivo := ".json"

	fmt.Printf("Insira o nome do seu personagem :%s", "\n")

	// input nos retorna a string e um erro
	Nome, _ = input.ReadString('\n')

	fmt.Printf("Insira a classe do seu personagem :%s", "\n")
	Classe, _ = input.ReadString('\n')

	// aqui estamos separando a string em uma lista de string
	// usando como parametro o espaco " "
	arch := strings.Split(Nome, " ")

	// arquivo nos retorna o resultado da concatenacao entre
	// arch[0], o primeiro nome digitado pelo usuario
	// + o conteudo de arquivo ".json"
	arquivo = arch[0] + arquivo

	holder := jogador{

		// Aqui estamos excluindo 2 caracteres do final da nossa string
		// que sao os caracteres "/r/n"
		// Responsaveis por encerrar uma linha
		Nome:   Nome[:len(Nome)-2],
		Classe: Classe[:len(Classe)-2],
		Level:  1,
	}

	// Marshal transforma struct em json
	file, _ := json.Marshal(holder)

	// ioutil cria o arquivo : nome, arquivo, configuracoes de seguranca
	ioutil.WriteFile(arquivo, file, 0644)

}

func main() {

	CriarJogador()
}
