package main

import "fmt"

func main() {

	v1 := [5]int{0, 1, 2, 3, 4}
	s1 := v1[:2]

	fmt.Println(v1, s1)

	s1[0] += 1
	fmt.Println(v1)

	m3 := make(map[int]string, 0)
	m3[0] := "eu robo"

	fmt.Print(m3)

	m2 := map[int]map[string]int{
		1: {"felipe": 30,
			"joao": 40},
		2: {"jose ": 80,
			"enrique": 120}}

	for _, valor := range m2 {
		fmt.Println(valor)

		for nome, _ := range valor {

			fmt.Print(nome)

		}
	}

}
