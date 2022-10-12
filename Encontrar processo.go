package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	// Biblioteca Process list
	// ela nos retorna todos os processos ativos
	ps "github.com/mitchellh/go-ps"
)

func Desligar() {

	// Comandos a serem executados no CMD
	arg0 := "shutdown"
	arg1 := "-s"
	arg2 := "-t"
	arg3 := "1"

	// Variavel que carrega cmd e comando a ser passado para ele
	cmd := exec.Command(arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print(string(stdout))

}

func AcharProcesso() bool {

	// Retorna todos os processos
	ListaProcessos, err := ps.Processes()
	if err != nil {
		log.Println("Erro em ps.Process")
		return false
	}

	// Passa por todos os processos
	// procurando o processo indicado na linha 52
	for x := range ListaProcessos {

		var process ps.Process
		process = ListaProcessos[x]

		a := fmt.Sprintf("%d\t%s\n", process.Pid(), process.Executable())

		if strings.Contains(a, "LeagueClient.exe") {
			Desligar()
			return true
		}
	}
	return false
}

func main() {

	// O loop e feito fora da funcao
	// pois se fosse feito dentro os valores nao atualizavam
	for AcharProcesso() == false {
	}

}
