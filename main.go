package main

import (
	"fmt"
	"sort"
)

func main() {
	var arreglo []string
	var auxString string
	var temp int

	fmt.Print("Cuantas palabras vas a ingresar? ")
	fmt.Scan(&temp)

	for temp > 0 {
		fmt.Print("Ingresa Palabra: ")
		fmt.Scan(&auxString)
		arreglo = append(arreglo, auxString)
		temp --
	}

	fmt.Println(arreglo)

	sort.Strings(arreglo)

	fmt.Print(arreglo)
}